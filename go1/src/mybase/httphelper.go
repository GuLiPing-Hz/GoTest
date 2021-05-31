package mybase

import (
	"crypto/subtle"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"

	"crypto/md5"
	"encoding/json"
	"math"
	"sort"
	"strconv"
)

const (
	TimeFmt             = "2006/01/02 15:04:05.000" //毫秒保留3位有效数字
	TimeFmtDB           = "2006-01-02 15:04:05"     //写入数据库用的时间
	TimeFmtSeq          = "20060102150405"          //yyyyMMddHHmmss
	TimeFmtSeqHW        = "20060102150405000"
	DateFmtDB           = "2006-01-02" //写入数据库用的日期
	AesKey       string = "Aabc#123admin@12"
)

//Http 返回值中的code状态码
const (
	WGSuccess         = iota //成功
	WGFail                   //请求失败
	WGErrorParam             //参数非法
	WGErrorTime              //时间戳非法
	WGErrorSign              //签名错误
	WGErrorTip               //提示msg信息
	WGErrorNeedLogin         //需要重新登录
	WGErrorParse             //json格式解析出错
	WGErrorNet               //网络异常
	WGErrorDataBase          //数据库操作失败
	WGErrorNoReg             //尚未注册 10
	WGIPForbidden            //该IP禁止访问 11
	WGErrorBusy              //客户端提示请勿频繁操作 12
	WGErrorRegistered        //已注册 13
	WGErrorNoImp             //未实现 14
	WGErrorEmpty             //当前库存为0
	WGNeedKefu               //请求可能成功，但是需要客服进一步处理
	WGSuccessWithTip         //成功处理，但是还需要异步回调通知。客户端可以先给提示。
	WGErrorExtBegin   = 1000 //扩展错误码起始

)

type HttpResult struct {
	Code int         `json:"code"` //状态码
	Msg  string      `json:"msg"`  //信息
	Data interface{} `json:"data"` //数据结构
}

func UrlEncode(param string) string {
	v := url.Values{}
	v.Add("encode", param)
	encoded := v.Encode()
	return encoded[strings.Index(encoded, "=")+1:]
}

func BuildResultEx(w http.ResponseWriter, r *http.Request, result string) {
	//if r != nil && r.Method == http.MethodOptions { //支持跨域访问
	//解决egret跨域访问的问题
	w.Header().Set("Access-Control-Allow-Origin", "*")
	//解决egret跨域访问head里面可以增加一些参数
	w.Header().Add("Access-Control-Allow-Headers", "curtime")
	w.Header().Add("Access-Control-Allow-Headers", "nonce")
	//}
	//fmt.Println("buildResult ", string(bs))
	_, _ = fmt.Fprint(w, result)
}

//构造http返回结果
func BuildResult(w http.ResponseWriter, r *http.Request, code int32, msg string, data interface{}) {
	var result = map[string]interface{}{}
	result["status"] = code //为了兼容老版本，后续会逐步去掉这个status，以code返回值为准
	result["code"] = code
	result["msg"] = msg
	if data != nil {
		result["data"] = data
	}

	bs, _ := json.Marshal(result)
	BuildResultEx(w, r, string(bs))
}

func BuildResult2(w http.ResponseWriter, r *http.Request, code int32, msg string) {
	BuildResult(w, r, code, msg, nil)
}

func BuildResult1(w http.ResponseWriter, r *http.Request, code int32) {
	BuildResult2(w, r, code, "")
}

func HttpGet(host, method string, param map[string]string, preFix, subFix string) (string, error) {
	return HttpGetEx(host, method, param, preFix, subFix, true)
}

func HttpGetEx(host, method string, param map[string]string, preFix, subFix string, needSign bool) (string, error) {
	return HttpGetUrl(fmt.Sprintf("%s/%s", host, method), param, preFix, subFix, needSign)
}

func HttpGetUrlNoSign(url string, param map[string]string) (string, error) {
	return HttpGetUrl(url, param, "", "", false)
}

func HttpGetUrl(httpUrl string, param map[string]string, preFix, subFix string, needSign bool) (string, error) {
	return HttpGetUrlEx(httpUrl, param, preFix, subFix, needSign, false)
}

func HttpGetUrlEx(httpUrl string, param map[string]string, preFix, subFix string, needSign, needJsonHead bool) (string, error) {
	values := url.Values{}
	if param != nil {
		for k, v := range param {
			values.Add(k, v)
		}
	}

	var sign = ""
	var curTimeStr = ""
	var nonce = ""
	if needSign {
		var paramsMap = make(map[string]string)
		curTime := time.Now().Unix()
		curTimeStr = strconv.FormatInt(curTime, 10)
		paramsMap["curtime"] = curTimeStr

		nonce = GetRandomString(6)
		paramsMap["nonce"] = nonce
		var paramKeys []string
		paramKeys = append(paramKeys, "nonce")
		paramKeys = append(paramKeys, "curtime")
		for k, _ := range param {
			paramsMap[k] = param[k]
			paramKeys = append(paramKeys, k)
		}
		sort.Strings(paramKeys)

		var paramSlice []string
		for i := range paramKeys {
			paramSlice = append(paramSlice, paramKeys[i]+"="+paramsMap[paramKeys[i]])
		}
		plainText := preFix + strings.Join(paramSlice, "&") + subFix
		sign = fmt.Sprintf("%x", md5.Sum([]byte(plainText)))
		values.Add("sign", sign)
	}

	urlFull := httpUrl
	if len(values) > 0 {
		urlFull = urlFull + "?" + values.Encode()
	}

	//提交请求
	reqest, err := http.NewRequest("GET", urlFull, nil)
	if err != nil {
		E("urlFull=%s,err=%s", urlFull, err)
		return "", err
	}

	if needSign {
		//增加header选项
		reqest.Header.Add("curtime", curTimeStr)
		reqest.Header.Add("nonce", nonce)
	}

	if needJsonHead {
		reqest.Header.Set("Accept", "application/json;charset=UTF-8")
	}

	//处理返回结果
	response, err := http.DefaultClient.Do(reqest)
	if err != nil {
		E("urlFull=%s,err=%s", urlFull, err)
		return "", err
	}
	bs, err := ioutil.ReadAll(response.Body)
	_ = response.Body.Close()
	if err != nil {
		E("urlFull=%s,err=%s", urlFull, err)
		return "", err
	}

	if response.StatusCode != 200 {
		W("http(%d) urlFull=%s", response.StatusCode, urlFull)
	}

	result := string(bs)
	fmt.Printf("url=[%s],head=%+v result=[%s]\n", urlFull, reqest.Header, result) //只在控制台打印一下。
	return result, nil
}

func HttpPostJson(strURL string, params, heads map[string]string) (string, error) {
	if heads == nil {
		heads = make(map[string]string)
	}
	heads["Content-Type"] = "application/json"

	theBody, err := json.Marshal(params)
	if err != nil {
		return "", err
	}
	return HttpPost(strURL, string(theBody), heads)
}

func HttpPostForm(strURL string, params, heads map[string]string) (string, error) {
	if heads == nil {
		heads = make(map[string]string)
	}
	heads["Content-Type"] = "application/x-www-form-urlencoded"

	values := url.Values{}
	for k, v := range params {
		values.Add(k, v)
	}

	return HttpPost(strURL, values.Encode(), heads)
}

func HttpPost(strURL, body string, heads map[string]string) (string, error) {
	fmt.Printf("SendPostJson url=%s\nbody=%v\n", strURL, body)
	req, err := http.NewRequest("POST", strURL, strings.NewReader(body))
	if err != nil {
		fmt.Printf("http.NewRequest is error:%s\n", err.Error())
		return "", err
	}

	if heads != nil {
		for k, v := range heads {
			req.Header.Add(k, v)
		}
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		E("http.NewRequest is error:%s.", err.Error())
		return "", err
	}
	respBodyBytes, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		E("urlFull=%s,body=%s,heads=%v,err=%s", strURL, body, heads, err)
		return "", err
	}
	return string(respBodyBytes), nil
}

/**
检查IP是否在白名单中
*/
func CheckIp(ip string, whiteIps []string) bool {
	for _, v := range whiteIps {
		if ip == v {
			return true
		}
	}
	return false
}

func SortParam(param map[string]string) string {
	var paramKeys []string
	for k := range param {
		paramKeys = append(paramKeys, k)
	}
	sort.Strings(paramKeys)

	var paramSlice []string
	for i := range paramKeys { //+的效率最高，参见 httphelper_test.go
		paramSlice = append(paramSlice, paramKeys[i]+"="+param[paramKeys[i]])
	}
	return strings.Join(paramSlice, "&")
}

func CheckHttpHeader(w http.ResponseWriter, r *http.Request, isProduct bool, preFix, subFix string) (bool, string) {
	_ = r.ParseForm()

	//I("CheckHttpHeader path=%s,form:%+v", r.URL.Path, r.Form)

	debug := r.Form.Get("debug")
	if !isProduct && debug == "1" {
		return true, ""
	}

	if r.Method == http.MethodOptions { //如果是跨域访问，直接返回成功
		fmt.Println("CheckHttpHeader http.MethodOptions")
		BuildResult1(w, r, WGSuccess)
		return false, ""
	}

	//获取传递的所有参数
	var paramsMap = make(map[string]string)
	var paramKeys []string
	for k, v := range r.Form {
		if k == "sign" { //跳过签名字段
			continue
		}
		paramsMap[k] = v[0]
		paramKeys = append(paramKeys, k)
	}

	curTime := r.Header.Get("curtime")
	paramKeys = append(paramKeys, "curtime")
	paramsMap["curtime"] = curTime
	nonce := r.Header.Get("nonce")
	paramKeys = append(paramKeys, "nonce")
	paramsMap["nonce"] = nonce

	body := "" //var body string= ""
	bodyMd5 := ""
	bits, _ := ioutil.ReadAll(r.Body)
	_ = r.Body.Close()

	if len(bits) > 0 {
		body = string(bits)
		bodyMd5 = fmt.Sprintf("%x", md5.Sum(bits[:]))
	}

	if bodyMd5 != "" {
		paramKeys = append(paramKeys, "body")
		paramsMap["body"] = bodyMd5
	}
	//fmt.Println("CheckHttpHeader curTime=", curTime, "nonce=", nonce)

	sort.Strings(paramKeys)
	var plainText = ""
	for _, v := range paramKeys {
		plainText += v + "=" + paramsMap[v] + "&"
	}
	plainText = preFix + plainText[0:len(plainText)-1] + subFix
	var sign = fmt.Sprintf("%x", md5.Sum([]byte(plainText)))
	var signReq = strings.ToLower(r.Form.Get("sign"))

	//@注意：防止计时攻击 参见：https://coolshell.cn/articles/21003.html  sign!=signReq
	if subtle.ConstantTimeCompare([]byte(sign), []byte(signReq)) != 1 {
		//fmt.Printf("CheckHttpHeader md5[%s] %s==%s\n", plainText, sign, signReq)
		E("CheckHttpHeader md5[%s] %s==%s", plainText, sign, signReq)
		BuildResult1(w, r, WGErrorSign)
		return false, ""
	}

	i64CurTime, err := strconv.ParseInt(curTime, 10, 64)
	if err != nil { //时间格式错误
		E("CheckHttpHeader time[%s] err[%s] ", curTime, err.Error())
		BuildResult1(w, r, WGErrorParam)
		return false, ""
	}

	if math.Abs(float64(time.Now().Unix()-i64CurTime)) > float64(300) { //300秒 10分钟内，上下5分钟
		//超过有效期
		E("CheckHttpHeader time[%v] outdate", i64CurTime)
		BuildResult1(w, r, WGErrorTime)
		return false, ""
	}
	return true, body
}
