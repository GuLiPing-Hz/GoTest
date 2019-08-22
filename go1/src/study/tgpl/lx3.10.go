package main

import (
	"bytes"
	"fmt"
	"strings"
)

func main() {
	fmt.Printf("练习题3.10 %s\n", comma("123459"))
	fmt.Printf("练习题3.11 %s\n", comma3_11("-1234567890.1112223334"))
	fmt.Printf("练习题3.12 %v\n", comma3_12("1国,2", "1,2国"))
}

// comma inserts commas in a non-negative decimal integer string.
func comma(s string) string {
	buf := &bytes.Buffer{}
	n := len(s)
	for i := 1; i <= n; i++ {
		buf.WriteByte(s[i-1])
		if i != n && (n-i)%3 == 0 {
			buf.WriteByte(',')
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
