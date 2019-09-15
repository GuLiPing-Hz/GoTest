package main

import (
	"reflect"
	"strconv"
	"fmt"
	"net/http"
	"strings"
	"bytes"
	"regexp"
)

func populate(v reflect.Value, val, check string) error {
	switch v.Kind() {
	case reflect.String:
		//练习题12.12
		if check == "email" {
			matched, _ := regexp.MatchString(`\w+@[\w.]+`, val)
			if !matched {
				return fmt.Errorf("the param need email,give %s", val)
			}
		}
		v.SetString(val)
	case reflect.Bool:
		b, err := strconv.ParseBool(val)
		if err != nil {
			return err
		}
		v.SetBool(b)
	case reflect.Int:
		x, err := strconv.ParseInt(val, 10, 32)
		if err != nil {
			return err
		}
		v.SetInt(x)
	default:
		return fmt.Errorf("unsupport type %d", v.Kind())
	}
	return nil
}

type RValue struct {
	val   reflect.Value
	check string
}

func Unpack(r *http.Request, ptr interface{}) error {
	if err := r.ParseForm(); err != nil {
		return err
	}

	fields := make(map[string]RValue)
	v := reflect.ValueOf(ptr).Elem()
	for i := 0; i < v.NumField(); i++ {
		fieldT := v.Type().Field(i)
		name := fieldT.Tag.Get("http")
		if name == "" {
			name = strings.ToLower(fieldT.Name)
		}
		fields[name] = RValue{v.Field(i),
			fieldT.Tag.Get("check"),
		} //field不需要Elem
	}

	//for k, v := range fields {
	//	if err := populate(v, r.Form.Get(k)); err != nil {
	//		return err
	//	}
	//}

	for k, vs := range r.Form {
		f, ok := fields[k]
		if !ok {
			continue
		}

		for _, v := range vs {
			if f.val.Kind() == reflect.Slice {
				val := reflect.New(f.val.Type().Elem()).Elem()
				if err := populate(val, v, ""); err != nil {
					return err
				}
				f.val.Set(reflect.Append(f.val, val))
			} else {
				if err := populate(f.val, v, f.check); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

func hFReflect(w http.ResponseWriter, r *http.Request) {
	var data struct {
		Labels     []string `http:"l"`
		MaxResults int      `http:"max"`
		Exact      bool     `http:"x"`
		Email      string   `http:"e" check:"email"`
	}
	data.MaxResults = 10 // set default
	if err := Unpack(r, &data); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest) // 400
		return
	}

	fmt.Fprintf(w, "Search: %+v\n", data)
}

func getStr(name string, v reflect.Value) (string, error) {
	switch v.Kind() {
	case reflect.Slice:
		str := &strings.Builder{}
		for i := 0; i < v.Len(); i++ {
			s, err := getStr(name, v.Index(i))
			if err != nil {
				return "", err
			}

			if i > 0 {
				str.WriteRune('&')
			}
			str.WriteString(s)
		}
		return str.String(), nil
	case reflect.String:
		return fmt.Sprintf("%s=%s", name, v.String()), nil
	case reflect.Int:
		return fmt.Sprintf("%s=%d", name, v.Int()), nil
	case reflect.Bool:
		return fmt.Sprintf("%s=%t", name, v.Bool()), nil
	}
	return "", fmt.Errorf("unsupport type =%d", v.Kind())
}

func Pack(data interface{}) ([]byte, error) {
	v := reflect.ValueOf(data)
	buf := &bytes.Buffer{}
	for i := 0; i < v.NumField(); i++ {
		fieldT := v.Type().Field(i)
		fieldV := v.Field(i)
		name := fieldT.Tag.Get("http")
		if name == "" {
			name = strings.ToLower(fieldT.Name)
		}

		str, err := getStr(name, fieldV)
		if err != nil {
			return nil, err
		}

		if i > 0 {
			buf.WriteRune('&')
		}
		buf.WriteString(str)
	}

	return buf.Bytes(), nil
}

func lx12_11() {
	var data = struct {
		Labels     []string `http:"l"`
		MaxResults int      `http:"max"`
		Exact      bool     `http:"x"`
	}{
		[]string{"a", "b"},
		5,
		true,
	}
	buf, err := Pack(data)
	if err != nil {
		fmt.Printf("err=%s\n", err.Error())
	}

	fmt.Printf("param=[%s]\n", string(buf))
}

func showReflectFunc(x interface{}) {
	fmt.Printf("type = %T\n", x)

	t := reflect.TypeOf(x)
	v := reflect.ValueOf(x)
	fmt.Printf("t==t2 %t\n", t == v.Type())

	for i := 0; i < v.NumMethod(); i++ {
		methodT := t.Method(i)
		methodV := v.Method(i)
		fmt.Printf("(%T).%s=%s\n", x, methodT.Name, methodT.Type.String())

		//使用Value调用函数，默认自带了当前实例，我们只需传递形参即可
		results := methodV.Call([]reflect.Value{reflect.ValueOf("Jack")})
		for i, v := range results {
			fmt.Printf("call1 result[%d]=%v\n", i, v)
		}

		//使用Type调用函数，相当于调用c的静态函数，但是我们需要把实例传入，类似Lua的self机制
		results2 := methodT.Func.Call([]reflect.Value{v, reflect.ValueOf("Jack")})
		for i, v := range results2 {
			fmt.Printf("call2 result[%d]=%v\n", i, v)
		}
	}
}

type HelloW struct {
}

func (h *HelloW) Print(x string) string {
	return fmt.Sprintf("Hello,I'm %s", x)
}

/**
反射的练习以及实验已经到头了，最后值得注意的是反射最好不要去滥用，相对于普通的写法，
反射无法让编译器识别类型错误，也影响了自动化重构和分析工具的准确性，
而且反射在运行速度有劣势，要慢一到两个数量级，
这就决定了关键核心部分代码不能用反射去偷懒实现。
原话参见《the Go programming language》458页

最多用于测试的时候可以用一下。
 */
func main() {
	//showReflectFunc(time.Hour)
	showReflectFunc(&HelloW{})

	lx12_11()

	http.HandleFunc("/reflect", hFReflect)
	if err := http.ListenAndServe(":8000", nil); err != nil {
		fmt.Printf("err=%s\n", err)
	}
}
