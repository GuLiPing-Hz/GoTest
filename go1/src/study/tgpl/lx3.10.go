package main

import (
	"bytes"
	"fmt"
	"strings"
)

//练习题3.13
const (
	KB int64 = 1000
	MB       = 1000 * KB
	GB       = 1000 * MB
	TB       = 1000 * KB
	PB       = 1000 * TB
	EB       = 1000 * PB
	ZB       = 1000 * EB
	YB       = 1000 * ZB
)

func main() {

	fmt.Printf("练习题3.10 %s\n", comma("123459"))
	fmt.Printf("练习题3.11 %s\n", comma3_11("-1234567890.1112223334"))
	fmt.Printf("练习题3.12 %v\n", comma3_12("1国,2", "1,2国"))

	fmt.Printf("KB=%d,MB=%d,YB=%d,YB/EB=%d\n", KB, MB, PB, YB/EB)//YB已经无法表示，但是可以计算，编译的时候已经计算 256bit
}

// comma inserts commas in a non-negative decimal integer string.
func comma(s string) string {
	buf := &bytes.Buffer{}
	n := len(s)
	for i := 1; i <= n; i++ {
		buf.WriteByte(s[i-1])
		if i != n && (n-i)%3 == 0 {
			fmt.Fprintf(buf, ",")
		}
	}
	return buf.String()
}

func comma3_11(s string) string {
	buf := &bytes.Buffer{}

	pos := strings.Index(s, ".")
	n := len(s)
	if pos != -1 {
		n = pos
	}
	offset := 1
	if n > 0 && s[0] == '+' || s[0] == '-' {
		buf.WriteByte(s[0])
		n = n - 1
		offset = 0
	}

	for i := 1; i <= n; i++ {
		buf.WriteByte(s[i-offset])
		if i != n && (n-i)%3 == 0 {
			buf.WriteByte(',')
		}
	}

	if pos != -1 {
		buf.WriteByte('.')
		n = len(s)
		for i := 1; i < n-pos; i++ {
			buf.WriteByte(s[i+pos])
			if i != n-pos-1 && i%3 == 0 {
				buf.WriteByte(',')
			}
		}
	}

	return buf.String()
}

func comma3_12(s1, s2 string) bool {
	for _, r := range s1 {
		if !strings.ContainsRune(s2, r) {
			return false
		}
	}

	for _, r := range s2 {
		if !strings.ContainsRune(s1, r) {
			return false
		}
	}
	return true
}
