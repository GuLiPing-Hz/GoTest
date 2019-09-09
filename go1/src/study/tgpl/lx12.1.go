package main

import (
	"fmt"
	"reflect"
)

func formatAtom(v reflect.Value) string {
	switch v.Kind() {
	case reflect.Bool:
		if v.Bool() {
			return "True"
		} else {
			return "False"
		}
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return fmt.Sprintf("%d", v.Int())
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return fmt.Sprintf("%d", v.Int())
	case reflect.String:
		return v.String()
	case reflect.Array:
		result := "["
		for i := 0; i < v.Len(); i++ {
			if i != 0 {
				result += ","
			}
			result = result + fmt.Sprintf("%s", formatAtom(v.Index(i)))
		}
		result += "]"
		return result
	case reflect.Struct:
		result := "{"
		for i := 0; i < v.NumField(); i++ {
			if i != 0 {
				result += ";"
			}
			result = result + fmt.Sprintf("%s", formatAtom(v.Field(i)))
		}
		result += "}"
		return result
	default:
		return "???"
	}
}

func display(path string, v reflect.Value, cur, limit int) {
	if cur >= limit {
		fmt.Println("out of limit")
		return
	}

	switch v.Kind() {
	case reflect.Invalid:
		fmt.Println(path, "Invalid")
	case reflect.Array, reflect.Slice:
		for i := 0; i < v.Len(); i++ {
			v2 := v.Index(i)
			display(fmt.Sprintf("%s[%d]", path, i), v2, cur+1, limit)
		}
	case reflect.Map:
		for _, k := range v.MapKeys() {
			fmt.Println(fmt.Sprintf("%s[%s]", path, formatAtom(k)), v.MapIndex(k))
		}
	case reflect.Struct:
		for i := 0; i < v.NumField(); i++ {
			display(fmt.Sprintf("%s.%s", path, v.Type().Field(i).Name), v.Field(i), cur+1, limit)
		}
	case reflect.Ptr:
		if v.IsNil() {
			fmt.Println(path, "nil")
		} else {
			display(fmt.Sprintf("* %s", path), v.Elem(), cur+1, limit)
		}
	case reflect.Interface:
		if v.IsNil() {
			fmt.Println(path, "nil")
		} else {
			display(fmt.Sprintf("%s.(%s)", path, v.Elem().Type().Name()), v.Elem(), cur+1, limit)
		}
	default:
		fmt.Println(path, formatAtom(v))
	}
}

func Display(path string, a interface{}) {
	display(path, reflect.ValueOf(a), 0, 50)
}

func main() {
	a := 10
	Display("&a", &a)

	fmt.Println()
	a1 := []struct {
		name  string
		score int
	}{{"Jack", 60}, {"Rose", 61}}
	Display("&a1", &a1)

	fmt.Println()
	a2 := make(map[string]int)
	a2["Jack"] = 62
	a2["Rose"] = 63
	Display("*a2", &a2)

	fmt.Println()
	a3 := make(map[[2]int]int)
	a3[[2]int{1, 2}] = 62
	Display("*a3", &a3)

	fmt.Println()
	type Person struct {
		name string
	}
	a4 := make(map[Person]int)
	a4[Person{"Jack"}] = 62
	Display("*a4", &a4)

	fmt.Println()
	type Cycle struct {
		val  int
		next *Cycle
	}
	a5 := Cycle{10, nil}
	a5.next = &a5
	Display("*a5", &a5)
}
