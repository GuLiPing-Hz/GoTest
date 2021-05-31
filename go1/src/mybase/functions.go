package mybase

import (
	"bytes"
	"crypto/hmac"
	"crypto/md5"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/mitchellh/mapstructure"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
	"io/ioutil"
	"math"
	"math/big"
	"net"
	"net/http"
	"os"
	"path/filepath"
	"runtime/debug"
	"strconv"
	"strings"
	"sync"
	"time"
)

var syncId sync.Mutex
var idSuffix int

// MD5 生成32位MD5,小写
func MD5(text string) string {
	ctx := md5.New()
	ctx.Write([]byte(text))
	return hex.EncodeToString(ctx.Sum(nil))
}

func HMACSHA256(text, key []byte) string {
	ctx := hmac.New(sha256.New, key)
	ctx.Write(text)
	return hex.EncodeToString(ctx.Sum(nil))
}

// maxValue > 0
func GetRandom(maxValue int) int {
	return RandInt(0, maxValue)
}

func RandInt(minValue, maxValue int) int {
	diff := maxValue - minValue
	ret, err := rand.Int(rand.Reader, big.NewInt(int64(diff)))
	if err != nil {
		return 0
	}
	return int(ret.Int64()) + minValue
}

// Abs 获取绝度值
func Abs(v int) int {
	return int(AbsI64(int64(v)))
}

func AbsI64(a int64) int64 {
	return int64(math.Abs(float64(a)))
}

// GenID 生成唯一ID
func GenID() string {
	syncId.Lock()
	defer syncId.Unlock()
	var id = Rand()
	return id.Hex()
}

func GenIDByTime() string {
	syncId.Lock()
	defer syncId.Unlock()
	idSuffix++
	if idSuffix > 255 {
		idSuffix = 0
	}
	var id string
	id = fmt.Sprintf("%d%d", time.Now().Unix(), idSuffix)
	return id
}

//// GetString 获取map 指定key对应的字符串值
//func GetString(d netdata.NetData, key string) (string, bool) {
//	return (&d).GetString(key)
//}
//
//func GetFloat64(d netdata.NetData, key string) (float64, bool) {
//	return (&d).GetFloat64(key)
//}
//
//func GetFloat32(d netdata.NetData, key string) (float32, bool) {
//	ret, ok := GetFloat64(d, key)
//	return float32(ret), ok
//}
//
//// GetString 获取map 指定key对应的字符串值
//func GetInt(d netdata.NetData, key string) (int, bool) {
//	ret, ok := GetInt32(d, key)
//	return int(ret), ok
//}
//
//func GetInt32(d netdata.NetData, key string) (int32, bool) {
//	ret, ok := GetInt64(d, key)
//	return int32(ret), ok
//}
//
//// GetString 获取map 指定key对应的字符串值
//func GetInt64(d netdata.NetData, key string) (int64, bool) {
//	return (&d).GetInt64(key)
//}
//
////从map中获取到bool类型的值
//func GetBool(d netdata.NetData, key string) (bool, bool) {
//	return (&d).GetBool(key)
//}
//
////从map类型中获取到slice的值
//func GetSlice(d map[string]interface{}, k string) ([]int64, bool) {
//	var plist []int64 = make([]int64, 0)
//	var gid int64
//	var v interface{}
//	var ok bool
//	if v, ok = d[k]; ok {
//		switch t := v.(type) {
//		case []interface{}:
//			length := len(v.([]interface{}))
//			for i := 0; i < length; i++ {
//				strgid := fmt.Sprint(v.([]interface{})[i])
//				gid, _ = strconv.ParseInt(strgid, 10, 64)
//				plist = append(plist, gid)
//			}
//			break
//		default:
//			_ = t
//		}
//	}
//	return plist, ok
//}
//
//从interface中获取对应的值并返回int值
//func Get(v interface{}) int {
//	var ret int
//	switch v.(type) {
//	case int:
//		ret = v.(int)
//		break
//	case string:
//		ret, _ = strconv.Atoi(v.(string))
//		break
//	case float64:
//		ret = int(v.(float64))
//		break
//	default:
//		return 0
//	}
//	return ret
//}
//
//func GetSlice2(d map[string]interface{}, k string) map[int]int {
//	var mapPlatRate map[int]int = make(map[int]int, 0)
//	if s, ok := d[k]; ok {
//		typer := reflect.TypeOf(s).Kind()
//		if typer == reflect.Slice {
//			var v []interface{}
//			v = d[k].([]interface{})
//			if len(v) > 0 {
//				for _, t := range v {
//					typet := reflect.TypeOf(t).Kind()
//					if typet == reflect.Slice {
//						pr := t.([]interface{})
//						prk := Get(pr[0])
//						prv := Get(pr[1])
//						mapPlatRate[prk] = prv
//					}
//				}
//			}
//		}
//	}
//	return mapPlatRate
//}

func ConvertVersion(version string) int64 {
	var versum int64
	s := strings.Split(version, ".")
	var l = len(s)
	if l == 2 {
		v0, _ := strconv.Atoi(s[0])
		v1, _ := strconv.Atoi(s[1])
		versum = (int64(v0) << 48) | (int64(v1) << 32)
	} else if l == 3 {
		v0, _ := strconv.Atoi(s[0])
		v1, _ := strconv.Atoi(s[1])
		v2, _ := strconv.Atoi(s[2])
		versum = (int64(v0) << 48) | (int64(v1) << 32) | (int64(v2) << 16)
	} else if l == 4 {
		v0, _ := strconv.Atoi(s[0])
		v1, _ := strconv.Atoi(s[1])
		v2, _ := strconv.Atoi(s[2])
		v3, _ := strconv.Atoi(s[3])
		versum = (int64(v0) << 48) | (int64(v1) << 32) | (int64(v2) << 16) | (int64(v3))
	}

	return versum
}

//byte转16进制字符串
func ByteToHex(data []byte) string {
	buffer := new(bytes.Buffer)
	for _, b := range data {
		s := strconv.FormatInt(int64(b&0xff), 16)
		if len(s) == 1 {
			buffer.WriteString("0")
		}
		buffer.WriteString(s)
	}
	return buffer.String()
}

//ipRange - a structure that holds the start and end of a range of ip addresses
type ipRange struct {
	start net.IP
	end   net.IP
}

// inRange - check to see if a given ip address is within a range given
func inRange(r ipRange, ipAddress net.IP) bool {
	// strcmp type byte comparison
	if bytes.Compare(ipAddress, r.start) >= 0 && bytes.Compare(ipAddress, r.end) < 0 {
		return true
	}
	return false
}

var privateRanges = []ipRange{
	ipRange{
		start: net.ParseIP("10.0.0.0"),
		end:   net.ParseIP("10.255.255.255"),
	},
	ipRange{
		start: net.ParseIP("100.64.0.0"),
		end:   net.ParseIP("100.127.255.255"),
	},
	ipRange{
		start: net.ParseIP("172.16.0.0"),
		end:   net.ParseIP("172.31.255.255"),
	},
	ipRange{
		start: net.ParseIP("192.0.0.0"),
		end:   net.ParseIP("192.0.0.255"),
	},
	ipRange{
		start: net.ParseIP("192.168.0.0"),
		end:   net.ParseIP("192.168.255.255"),
	},
	ipRange{
		start: net.ParseIP("198.18.0.0"),
		end:   net.ParseIP("198.19.255.255"),
	},
}

// isPrivateSubnet - check to see if this ip is in a private subnet
func isPrivateSubnet(ipAddress net.IP) bool {
	// my use case is only concerned with ipv4 atm
	if ipCheck := ipAddress.To4(); ipCheck != nil {
		// iterate over all our ranges
		for _, r := range privateRanges {
			// check if this ip is in a private range
			if inRange(r, ipAddress) {
				return true
			}
		}
	}
	return false
}

func getIPAdress(r *http.Request) string {
	for _, h := range []string{"X-Forwarded-For", "X-Real-Ip"} {
		addresses := strings.Split(r.Header.Get(h), ",")
		// march from right to left until we get a public address
		// that will be the address right before our proxy.
		for i := len(addresses) - 1; i >= 0; i-- {
			__ip := strings.TrimSpace(addresses[i])
			// header can contain spaces too, strip those out.
			realIP := net.ParseIP(__ip)
			if !realIP.IsGlobalUnicast() || isPrivateSubnet(realIP) {
				// bad address, go to next
				continue
			}
			return __ip
		}
	}
	return ""
}

func EarthDistance(lat1, lng1, lat2, lng2 float64) float64 {
	radius := 6371000.0 // 6378137
	rad := math.Pi / 180.0

	lat1 = lat1 * rad
	lng1 = lng1 * rad
	lat2 = lat2 * rad
	lng2 = lng2 * rad

	theta := lng2 - lng1
	dist := math.Acos(math.Sin(lat1)*math.Sin(lat2) + math.Cos(lat1)*math.Cos(lat2)*math.Cos(theta))

	return dist * radius
}

func IsPhoneNumber(str string) bool {
	if len(str) != 11 {
		return false
	}

	for _, v := range str {
		if v >= '0' && v <= '9' {

		} else {
			return false
		}
	}
	return true
}

func IsPwd(str string) bool {
	if len(str) >= 64 {
		return false
	}

	for _, v := range str {
		if (v >= '0' && v <= '9') || (v >= 'a' && v <= 'f') || (v >= 'A' && v <= 'F') {

		} else {
			return false
		}
	}
	return true
}

func IsUname(str string) bool {
	if len(str) < 6 || len(str) >= 64 {
		return false
	}
	return true
}

func IsCode(str string) bool {
	if len(str) != 4 {
		return false
	}

	for _, v := range str {
		if (v >= '0' && v <= '9') ||
			(v >= 'a' && v <= 'z') ||
			(v >= 'A' && v <= 'Z') {

		} else {
			return false
		}
	}
	return true
}

func GenTokenRobot(uid int64) string {
	strSrc := fmt.Sprintf("%d-", uid)
	strMD5 := MD5(strSrc)
	return strMD5
}

func GenToken(uid int64) string {
	strRandom1 := GetRandomString(4)
	strSrc := fmt.Sprintf("%s-%d", strRandom1, uid)

	bbAES, err := AESEncrypt([]byte(strSrc), []byte(AesKey))
	if err != nil {
		return ""
	}

	buf := make([]byte, hex.EncodedLen(len(bbAES)))
	hex.Encode(buf, bbAES)
	return string(buf)
}

func GetRandomString(l int) string {
	bs := []byte("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	res := make([]byte, 0, l)
	for i := 0; i < l; i++ {
		n := GetRandom(len(bs))
		res = append(res, bs[n])
	}
	return string(res)
}

func GbkToUtf8(s []byte) ([]byte, error) {
	reader := transform.NewReader(bytes.NewReader(s), simplifiedchinese.GBK.NewDecoder())
	d, e := ioutil.ReadAll(reader)
	if e != nil {
		return nil, e
	}
	return d, nil
}

func Utf8ToGbk(s []byte) ([]byte, error) {
	reader := transform.NewReader(bytes.NewReader(s), simplifiedchinese.GBK.NewEncoder())
	d, e := ioutil.ReadAll(reader)
	if e != nil {
		return nil, e
	}
	return d, nil
}

func LoadCfg(filename string, cfg interface{}) error {
	filePath, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		E("path err=%v", err)
		return err
	}
	fullPathFile := filePath + "/" + filename
	buf, err := ioutil.ReadFile(fullPathFile)
	if err != nil {
		E("LoadCfg ReadFile[%s]: %s", fullPathFile, err.Error())
		return err
	}

	if err := json.Unmarshal(buf, cfg); err != nil {
		E("LoadCfg Unmarshal error[%s]: %s", fullPathFile, err.Error())
		return err
	}

	return nil
}

//获取到指定时间的0点的time.Time
func GetTodayMidnightEx(theTime *time.Time) time.Time {
	if theTime == nil {
		now := time.Now()
		theTime = &now
	}
	strTime := fmt.Sprintf("%04d-%02d-%02d 00:00:00", theTime.Year(), theTime.Month(), theTime.Day())
	midnight, err := time.Parse(TimeFmtDB, strTime)
	if err != nil {
		E("err=%v", err)
	}
	//fmt.Printf("midnight = %d\n", midnight.Unix())
	return midnight
}

//获取到今天0点的time.Time
func GetTodayMidnight() time.Time {
	return GetTodayMidnightEx(nil)
}

/**
map[string]interface{} ->数据结构
数据结构 -> map[string]interface{}

@param input []map[string]interface 或者 map[string]interface 或者 结构
@param output 结构指针 或者 map指针

@return nil无错误
*/
func Decode(input, outputPtr interface{}) error {
	return DecodeEx(input, outputPtr, false)
}

/**
map[string]string ->数据结构
*/
func DecodeRedis(input, outputPtr interface{}) error {
	return DecodeEx(input, outputPtr, true)
}

/**
@outputPtr 需要指针类型
*/
func DecodeEx(input, outputPtr interface{}, weakly bool) error {
	//dataType := reflect.TypeOf(outputPtr) //获取数据类型
	//if dataType.Kind() != reflect.Ptr {
	//	return fmt.Errorf("need Ptr")
	//}
	config := &mapstructure.DecoderConfig{
		Metadata:         nil,
		Result:           outputPtr,
		TagName:          "json",
		WeaklyTypedInput: weakly,
	}
	decoder, err := mapstructure.NewDecoder(config)
	if err != nil {
		return err
	}

	return decoder.Decode(input)
}

//捕获异常，并写入到指定文件中。
func PanicHandler() {
	if err := recover(); err != nil {
		exeName := os.Args[0] //获取程序名称

		now := time.Now() //获取当前时间

		time_str := now.Format("20060102150405")                  //设定时间格式
		fname := fmt.Sprintf("%s-%s-dump.log", exeName, time_str) //保存错误信息文件名:程序名-进程ID-当前时间（年月日时分秒）

		f, err1 := os.Create(fname)
		if err1 != nil {
			return
		}
		defer f.Close()
		builder := &strings.Builder{}
		builder.Write(debug.Stack())
		fmt.Printf("dump to file err=%s\nstack=\n%s", err, builder.String())
		_, _ = f.WriteString(fmt.Sprintf("%v\r\n", err)) //输出panic信息
		_, _ = f.WriteString("========\r\n")
		_, _ = f.WriteString(string(debug.Stack())) //输出堆栈信息
	}
}

func GetRandSeed() int64 {
	var a = 0 //变量地址当做随机数
	var b = 0 //变量地址当做随机数
	aPtr, _ := strconv.ParseInt(fmt.Sprintf("%x", &a), 16, 64)
	bPtr, _ := strconv.ParseInt(fmt.Sprintf("%x", &b), 16, 64)

	return time.Now().Unix() * aPtr * bPtr
}

//各个服务器业务相关的一些通用方法
func GetUsrKeyInRedis(uid int64) string {
	return fmt.Sprintf("usr_%d", uid)
}

func GetUsrFlagInRedis(uid int64) string {
	return fmt.Sprintf("usr_key_%d", uid)
}

func GetUsrListKeyInRedis() string {
	return "redis_usrs" //"usr_redis"
}

func GetUsrTokenInRedis(uid int64) string {
	return fmt.Sprintf("token_%d", uid)
}

func GetUsrCheckInKeyInRedis(uid int64) string {
	return fmt.Sprintf("usr_checkin_%d", uid)
}

func GetUsrCheckIn7KeyInRedis(uid int64) string {
	return fmt.Sprintf("usr_checkin7_%d", uid)
}

//仅供查看数据
func GetUsrTodayInRedis(uid int64) string {
	return fmt.Sprintf("usr_today_%d", uid)
}
func GetUsrTodayMInRedis(uid int64) string {
	return fmt.Sprintf("usr_today_m_%d", uid)
}
