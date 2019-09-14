package main

import (
	"reflect"
	"strconv"
	"fmt"
	"net/http"
	"strings"
)

func populate(v reflect.Value, val string) error {
	switch v.Kind() {
	case reflect.String:
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

func Unpack(r *http.Request, ptr interface{}) error {
	if err := r.ParseForm(); err != nil {
		return err
	}

	fields := make(map[string]reflect.Value)
	v := reflect.ValueOf(ptr).Elem()
	for i := 0; i < v.NumField(); i++ {
		field := v.Type().Field(i)
		name := field.Tag.Get("http")
		if name == "" {
			name = strings.ToLower(field.Name)
		}
		fields[name] = v.Field(i) //field不需要Elem
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
			if f.Kind() == reflect.Slice {
				val := reflect.New(f.Type().Elem()).Elem()
				if err := populate(val, v); err != nil {
					return err
				}
				f.Set(reflect.Append(f, val))
			} else {
				if err := populate(f, v); err != nil {
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
	}
	data.MaxResults = 10 // set default
	if err := Unpack(r, &data); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest) // 400
		return
	}

	fmt.Fprintf(w, "Search: %+v\n", data)
}

func main() {
	http.HandleFunc("/reflect", hFReflect)
	if err := http.ListenAndServe(":8000", nil); err != nil {
		fmt.Printf("err=%s\n", err)
	}
}
