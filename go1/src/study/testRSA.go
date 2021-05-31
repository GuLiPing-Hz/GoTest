package main

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"fmt"
	"go1/src/mybase"
	"golang.org/x/crypto/pkcs12"
	"io/ioutil"
	"net/url"
	"sort"
	"strings"
	"time"
)

//私钥签名
func RsaSign(privateKey, data []byte) ([]byte, error) {
	h := sha256.New()
	h.Write(data)
	hashed := h.Sum(nil)
	//获取私钥
	block, _ := pem.Decode(privateKey)
	if block == nil {
		return nil, fmt.Errorf("private key invalid")
	}
	//解析PKCS1格式的私钥
	priv, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	return rsa.SignPKCS1v15(rand.Reader, priv, crypto.SHA256, hashed)
}

//公钥验证
func RsaSignVer(publicKey, data []byte, signature []byte) error {
	hashed := sha256.Sum256(data)
	block, _ := pem.Decode(publicKey)
	if block == nil {
		return fmt.Errorf("public key invalid")
	}
	// 解析公钥
	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return err
	}
	// 类型断言
	pub := pubInterface.(*rsa.PublicKey)
	//验证签名
	return rsa.VerifyPKCS1v15(pub, crypto.SHA256, hashed[:], signature)
}

// 公钥加密
func RsaEncrypt(publicKey, data []byte) ([]byte, error) {
	//解密pem格式的公钥
	block, _ := pem.Decode(publicKey)
	if block == nil {
		return nil, fmt.Errorf("public key invalid")
	}
	// 解析公钥
	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	// 类型断言
	pub := pubInterface.(*rsa.PublicKey)
	//加密
	return rsa.EncryptPKCS1v15(rand.Reader, pub, data)
}

// 私钥解密
func RsaDecrypt(privateKey, ciphertext []byte) ([]byte, error) {
	//获取私钥
	block, _ := pem.Decode(privateKey)
	if block == nil {
		return nil, fmt.Errorf("private key invalid")
	}
	//解析PKCS1格式的私钥
	priv, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	// 解密
	return rsa.DecryptPKCS1v15(rand.Reader, priv, ciphertext)
}

// 根据银联获取到的PFX文件和密码来解析出里面包含的私钥(rsa)和证书(x509)
func ParserPfxToCert(path string, password string) (private *rsa.PrivateKey, cert *x509.Certificate, err error) {
	var pfxData []byte
	pfxData, err = ioutil.ReadFile(path)
	if err != nil {
		return nil, nil, err
	}
	priv, cert, err := pkcs12.Decode(pfxData, password)
	if err != nil {
		return nil, nil, err
	}
	private = priv.(*rsa.PrivateKey)
	return private, cert, nil
}

// 根据文件名解析出证书
func ParseCertificateFromFile(path string) (cert *x509.Certificate, err error) {
	// Read the verify sign certification key
	pemData, err := ioutil.ReadFile(path)
	if err != nil {
		return
	}

	// Extract the PEM-encoded data block
	block, _ := pem.Decode(pemData)
	if block == nil {
		err = fmt.Errorf("bad key data: %s", "not PEM-encoded")
		return
	}
	if got, want := block.Type, "CERTIFICATE"; got != want {
		err = fmt.Errorf("unknown key type %q, want %q", got, want)
		return
	}

	// Decode the certification
	cert, err = x509.ParseCertificate(block.Bytes)
	if err != nil {
		err = fmt.Errorf("bad private key: %s", err)
		return
	}
	return
}

// base64 加密
func base64Encode(data []byte) string {
	return base64.StdEncoding.EncodeToString(data)
}

// base64 解密
func base64Decode(data string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(data)
}

func SortDataForHttpResp(param url.Values, excludes map[string]bool) string {
	var paramKeys []string
	for k, _ := range param {
		if excludes != nil {
			if _, ok := excludes[k]; ok {
				continue
			}
		}
		paramKeys = append(paramKeys, k)
	}
	sort.Strings(paramKeys)
	var plainText = ""
	for _, v := range paramKeys {
		plainText += v + "=" + param.Get(v) + "&"
	}
	return plainText[0 : len(plainText)-1]
}

func ylCheckSign(rootCer, middleCer *x509.Certificate, resData url.Values) error {
	var signPubKeyCert = resData.Get("signPubKeyCert")

	//1.从返回报文中获取公钥信息转换成公钥对象
	//String strCert = resData.get(SDKConstants.param_signPubKeyCert);
	mybase.I("验签公钥证书：[" + signPubKeyCert + "]")

	//X509Certificate x509Cert = CertUtil.genCertificateByStr(strCert);
	//if(x509Cert == null) {
	//	LogUtil.writeErrorLog("convert signPubKeyCert failed");
	//	return false;
	//}
	// Extract the PEM-encoded data block
	var pub *rsa.PublicKey
	//pubInterface, ok := MapKeyPubs.Load(signPubKeyCert)
	//if !ok {
	block, _ := pem.Decode([]byte(signPubKeyCert))
	if block == nil {
		return fmt.Errorf("bad key data=%s", "not PEM-encoded")
	}
	if "CERTIFICATE" != block.Type {
		return fmt.Errorf("unknown key type=%s, want CERTIFICATE", block.Type)
	}
	x509Cert, err := x509.ParseCertificate([]byte(signPubKeyCert))
	if err != nil {
		return err
	}

	// 2.验证证书链
	//if (!CertUtil.verifyCertificate(x509Cert)) {
	//	LogUtil.writeErrorLog("验证公钥证书失败，证书信息：[" + strCert + "]");
	//	return false;
	//}
	//检查有效期有没有过期，启用时间没必要检查了。
	now := time.Now()
	if now.After(x509Cert.NotAfter) {
		return fmt.Errorf("certificate expired")
	}

	//验证证书链，
	roots := x509.NewCertPool()
	roots.AddCert(rootCer)
	intermediateCerts := x509.NewCertPool()
	intermediateCerts.AddCert(middleCer)
	intermediateCerts.AddCert(rootCer)
	//链式向上验证证书
	//验证用户证书
	opts := x509.VerifyOptions{
		KeyUsages:     []x509.ExtKeyUsage{x509.ExtKeyUsageAny},
		Intermediates: intermediateCerts,
		Roots:         roots,
	}
	if _, err := x509Cert.Verify(opts); err != nil {
		return fmt.Errorf("failed to verify certificate Chain::" + err.Error())
	}

	//验证证书是否属于银联
	cn := strings.Split(x509Cert.Subject.CommonName, "@")[2]
	if "中国银联股份有限公司" != cn {
		return fmt.Errorf("证书不属于银联")
	}

	pub = x509Cert.PublicKey.(*rsa.PublicKey)
	//	MapKeyPubs.Store(signPubKeyCert, pub)
	//} else {
	//	pub = pubInterface.(*rsa.PublicKey)
	//}

	var signature = resData.Get("signature")
	var exclueds = make(map[string]bool)
	exclueds["signature"] = true
	exclueds["signPubKeyCert"] = true
	plainText := SortDataForHttpResp(resData, exclueds)

	//hashed := fmt.Sprintf("%x", sha256.Sum256([]byte(plainText)))
	hashed := sha256.Sum256([]byte(plainText))
	inSign, err := base64.StdEncoding.DecodeString(signature)
	if err != nil {
		return fmt.Errorf("解析返回signature失败 err=%s", err)
	}
	err = rsa.VerifyPKCS1v15(pub, crypto.SHA256, hashed[:], inSign)
	if err != nil {
		mybase.I("签名原文：[%s] encoding=%s", signature, resData.Get("encoding"))
		mybase.I("待验签返回报文串：[" + string(hashed[:]) + "]")
		return fmt.Errorf("返回数据验签失败 err=%s", err.Error())
	}
	return nil
}

func SortDataForHttpReq(param map[string]string) string {
	var paramKeys []string
	for k, _ := range param {
		paramKeys = append(paramKeys, k)
	}
	sort.Strings(paramKeys)
	var plainText = ""
	for _, v := range paramKeys {
		plainText += v + "=" + param[v] + "&"
	}
	return plainText[0 : len(plainText)-1]
}

func ylSign(private *rsa.PrivateKey, param map[string]string) (string, error) {
	plainText := SortDataForHttpReq(param)

	hashed := sha256.Sum256([]byte(plainText))
	signer, err := rsa.SignPKCS1v15(rand.Reader, private, crypto.SHA256, hashed[:])
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(signer), nil
}
