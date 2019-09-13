package main

import (
	"reflect"
	"io"
	"fmt"
	"bytes"
	"text/scanner"
	"strconv"
)

func encode(w io.Writer, v reflect.Value) error {
	switch v.Kind() {
	case reflect.Invalid:
		io.WriteString(w, "nil")
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		fmt.Fprintf(w, "%d", v.Int())
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		fmt.Fprintf(w, "%d", v.Uint())
	case reflect.String:
		fmt.Fprintf(w, "%q", v.String())
	case reflect.Bool:
		if v.Bool() {
			io.WriteString(w, "True")
		} else {
			io.WriteString(w, "False")
		}
	case reflect.Float32, reflect.Float64:
		fmt.Fprintf(w, "%g", v.Float())
	case reflect.Complex64, reflect.Complex128:
		c := v.Complex()
		fmt.Fprintf(w, "#C(%f %f)", real(c), imag(c))
	case reflect.Ptr, reflect.Interface:
		return encode(w, v.Elem())
	case reflect.Array, reflect.Slice: //(val ...)
		fmt.Fprintf(w, "%s (", v.Type().Name())
		for i := 0; i < v.Len(); i++ {
			if i > 0 {
				io.WriteString(w, " ")
			}

			if err := encode(w, v.Index(i)); err != nil {
				return err
			}
		}
		io.WriteString(w, ")")
	case reflect.Struct: // ((name value) ...)
		fmt.Fprintf(w, "(")
		for i := 0; i < v.NumField(); i++ {
			if i > 0 {
				io.WriteString(w, " ")
			}

			io.WriteString(w, "(")
			io.WriteString(w, v.Type().Field(i).Name)
			io.WriteString(w, " ")
			if err := encode(w, v.Field(i)); err != nil {
				return err
			}
			io.WriteString(w, ")")
		}
		io.WriteString(w, ")")
	case reflect.Map: // ((key value) ...)
		io.WriteString(w, "(")
		for _, k := range v.MapKeys() {
			io.WriteString(w, "(")
			if err := encode(w, k); err != nil {
				return err
			}
			io.WriteString(w, " ")
			if err := encode(w, v.MapIndex(k)); err != nil {
				return err
			}
			io.WriteString(w, ")")
		}
		io.WriteString(w, ")")
	default: // float, complex, bool, chan, func, interface
		return fmt.Errorf("unsupported type: %s", v.Type())
	}
	return nil
}

func Marshal(v interface{}) ([]byte, error) {
	var buf = new(bytes.Buffer)
	if err := encode(buf, reflect.ValueOf(v)); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

type Lex struct {
	sc *scanner.Scanner
	ch rune
}

func (l *Lex) next() {
	l.ch = l.sc.Scan()
}

func (l *Lex) text() string {
	return l.sc.TokenText()
}

func (l *Lex) consume(c rune) error {
	if l.ch != c {
		return fmt.Errorf("give %c,want %c", l.ch, c)
	}
	l.next()
	return nil
}

func endList(l *Lex) bool {
	switch l.ch {
	case scanner.EOF:
		panic("end of file")
	case ')':
		return true
	}
	return false
}

func readList(l *Lex, v reflect.Value) error {
	switch v.Kind() {
	case reflect.Array: // (item ...)
		for i := 0; !endList(l); i++ {
			read(l, v.Index(i).Elem())
		}
	case reflect.Slice: // (item ...)
		for !endList(l) {
			item := reflect.New(v.Type().Elem()).Elem() //取出slice中单个元素的类型
			read(l, item)
			v.Set(reflect.Append(v, item))
		}
	case reflect.Map: // ((key value) ...)
		v.Set(reflect.MakeMap(v.Type()))
		for !endList(l) {
			if err := l.consume('('); err != nil {
				return err
			}
			key := reflect.New(v.Type().Key()).Elem() //取出map中的key值类型
			read(l, key)
			value := reflect.New(v.Type().Elem()).Elem() //取出map中的value值类型
			read(l, value)
			v.SetMapIndex(key, value)

			if err := l.consume(')'); err != nil {
				return err
			}
		}
	case reflect.Struct: // ((name value) ...)
		for !endList(l) {
			if err := l.consume('('); err != nil {
				return err
			}

			name := l.text()
			l.next()
			read(l, v.FieldByName(name)) //根据field名字找到对应的field

			if err := l.consume(')'); err != nil {
				return err
			}
		}
	default:
		return fmt.Errorf("unknow type %s\n", v.Kind().String())
	}

	return nil
}

func read(l *Lex, v reflect.Value) error {
	switch l.ch {
	case scanner.Ident: //类似关键字的意思。。
		if l.text() == "nil" {
			v.Set(reflect.Zero(v.Type()))
			l.next()
		} else if l.text() == "True" {
			v.SetBool(true)
			l.next()
		} else if l.text() == "False" {
			v.SetBool(false)
			l.next()
		}
	case scanner.String:                  //带引号的字符串
		s, _ := strconv.Unquote(l.text()) // NOTE: ignoring errors
		v.SetString(s)
		l.next()
	case scanner.Int:
		s, _ := strconv.ParseInt(l.text(), 10, 64) // NOTE: ignoring errors
		v.SetInt(s)
		l.next()
	case '(':
		l.next()
		readList(l, v)
		l.next()
	default:
		return fmt.Errorf("unknow ch %v\n", l.ch)
	}
	return nil
}

func Unmarshal(buf []byte, out interface{}) error {
	v := reflect.ValueOf(out).Elem()
	sc := scanner.Scanner{}
	sc.Init(bytes.NewReader(buf))

	lex := Lex{sc: &sc}
	lex.next()
	read(&lex, v)
	return nil
}

type Movie struct {
	Title, Subtitle string
	Year            int
	Color           bool
	Actor           map[string]string
	Oscars          []string
	Sequel          *string
}

func main() {
	strangelove := Movie{
		Title:    "Dr. Strangelove",
		Subtitle: "How I Learned to Stop Worrying and Love the Bomb",
		Year:     1964,
		Color:    false,
		Actor: map[string]string{
			"Dr. Strangelove":            "Peter Sellers",
			"Grp. Capt. Lionel Mandrake": "Peter Sellers",
			"Pres. Merkin Muffley":       "Peter Sellers",
			"Gen. Buck Turgidson":        "George C. Scott",
			"Brig. Gen. Jack D. Ripper":  "Sterling Hayden",
			`Maj. T.J. "King" Kong`:      "Slim Pickens",
		},
		Oscars: []string{
			"Best Actor (Nomin.)",
			"Best Adapted Screenplay (Nomin.)",
			"Best Director (Nomin.)",
			"Best Picture (Nomin.)",
		},
	}
	buf, err := Marshal(strangelove)
	if err != nil {
		fmt.Printf("err=%s\n", err.Error())
		return
	}
	fmt.Println(string(buf))

	a := 2
	b := reflect.ValueOf(a)
	c := reflect.ValueOf(&a) //单纯的指针类型的反射值是不可以寻址的，他只是保存了指针类型
	//通过Elem函数才能寻址，才能更新原来的变量
	//貌似结构中的field天然就是可寻址的反射变量
	d := reflect.ValueOf(&a).Elem()
	fmt.Println("是否可以寻址", b.CanAddr(), c.CanAddr(), d.CanAddr())

	//未导出的变量通常不能寻址
	fmt.Println("是否可以更新", b.CanSet(), c.CanSet(), d.CanSet())

	obj := new(Movie)
	err = Unmarshal(buf, obj)
	if err != nil {
		fmt.Printf("err=%s\n", err.Error())
	} else {
		fmt.Printf("obj=%v\n", *obj)
	}

}
