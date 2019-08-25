package main

import (
	"unicode/utf8"
	"unicode"
	"fmt"
)

func main() {
	fmt.Printf("r=%d", panic1())
}

//练习题4.3
func reverse(s *[]int) {
	for i, j := 0, len(*s)-1; i < j; i, j = i+1, j-1 {
		(*s)[i], (*s)[j] = (*s)[j], (*s)[i]
	}
}

//练习题4.6
func space(b *[]byte) *[]byte {
	b2 := b
	pos := 0
	for {
		r, size1 := utf8.DecodeRune((*b2)[pos:])
		if r == utf8.RuneError {
			break
		}
		pos += size1
		r2, size2 := utf8.DecodeRune((*b2)[pos:])
		if r2 == utf8.RuneError {
			break
		}
		pos2 := pos + size2
		if unicode.IsSpace(r) && unicode.IsSpace(r2) {
			copy((*b2)[pos:], (*b2)[pos2:])
			pos -= size1
		}
	}
	return b
}

//练习题4.7
func reverse2(b *[]byte) {
	for i, j := 0, len(*b)-1; i < j; i, j = i+1, j-1 {
		(*b)[i], (*b)[j] = (*b)[j], (*b)[i]
	}
}

//练习题 5.19
func panic1() (r int) {
	defer func() {
		p := recover()
		r = p.(int)
	}()

	panic(5)
}
