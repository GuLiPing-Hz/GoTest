package netdata

import (
	"reflect"
)

type NetData map[string]interface{}

func (this *NetData) GetInt64(key string) (ret int64, ok bool) {
	//if err := recover(); err != nil {
	//	ret = 0
	//	ok = false
	//}

	dataI, ok := (*this)[key]
	if !ok {
		return 0, false
	}

	rV := reflect.ValueOf(dataI)
	if rV.Kind() == reflect.Int || rV.Kind() == reflect.Int8 || rV.Kind() == reflect.Int16 ||
		rV.Kind() == reflect.Int32 || rV.Kind() == reflect.Int64 {
		return rV.Int(), true
	}
	retF, ok := this.GetFloat64(key)
	if ok {
		return int64(retF), true
	}
	return 0, false
}

func (this *NetData) GetInt32(key string) (ret int32, ok bool) {
	temp, ok := this.GetInt64(key)
	return int32(temp), ok
}

func (this *NetData) GetInt(key string) (int, bool) {
	ret, ok := this.GetInt64(key)
	return int(ret), ok
}

func (this *NetData) GetUInt64(key string) (uint64, bool) {
	dataI, ok := (*this)[key]
	if !ok {
		return 0, false
	}

	rV := reflect.ValueOf(dataI)
	if rV.Kind() == reflect.Uint || rV.Kind() == reflect.Uint8 || rV.Kind() == reflect.Uint16 ||
		rV.Kind() == reflect.Uint32 || rV.Kind() == reflect.Uint64 || rV.Kind() == reflect.Uintptr {
		return rV.Uint(), true
	}
	v, ok := this.GetInt64(key)
	if ok {
		return uint64(v), true
	}
	return 0, false
}

func (this *NetData) GetFloat64(key string) (float64, bool) {
	dataI, ok := (*this)[key]
	if !ok {
		return 0, false
	}

	rV := reflect.ValueOf(dataI)
	if rV.Kind() == reflect.Float32 || rV.Kind() == reflect.Float64 {
		return rV.Float(), true
	}
	return 0, false
}

func (this *NetData) GetString(key string) (string, bool) {
	dataI, ok := (*this)[key]
	if !ok {
		return "", false
	}

	rV := reflect.ValueOf(dataI)
	if rV.Kind() == reflect.String {
		return rV.String(), true
	}
	return "", false
}

func (this *NetData) GetBool(key string) (bool, bool) {
	dataI, ok := (*this)[key]
	if !ok {
		return false, false
	}

	rV := reflect.ValueOf(dataI)
	if rV.Kind() == reflect.Bool {
		return rV.Bool(), true
	}
	return false, false
}

func (this *NetData) GetInterface(key string) (interface{}, bool) {
	dataI, ok := (*this)[key]
	return dataI, ok
}

func NewData() *NetData {
	ret := make(NetData)
	return &ret
}
