package main

import (
	"bytes"
	"fmt"
	"golang.org/x/net/html"
	"os"
	"path/filepath"
	"pkg"
	"unicode"
	"unicode/utf8"
	"unsafe"
	"bufio"
	"io"
	"os"
)

func main() {
	fmt.Printf("r=%d\n", panic1())

	fmt.Printf("^uint(0)=%b,^uint(0)>>63=%b,32 << (^uint(0) >> 63)=%v;;;sizeof(uint)=%d\n",
		^uint(0), ^uint(0)>>63, 32<<(^uint(0)>>63), unsafe.Sizeof(uint(0))*8)

	bc := &ByteCounter{}
	bc.Write([]byte("Jack has an apple.\nNow he is eating the apple."))
	fmt.Printf("bc=%v\n", bc)

	w, n := CountingWriter(os.Stdout)
	w.Write([]byte("hello 111"))
	fmt.Printf("\nn=%d\n", *n)

	if len(os.Args) < 2 {
		fmt.Printf("need file\n")
		return
	}

	fileDir := filepath.Dir(os.Args[0])
	filePath := fmt.Sprintf("%s%c%s", fileDir, filepath.Separator, os.Args[1])
	file, err := os.Open(filePath)

	if err != nil {
		fmt.Printf("err = %s\n", err.Error())
		return
	}
	doc, err := html.Parse(file)
	if err != nil {
		fmt.Printf("err=%s\n", err.Error())
		return
	}

	nameCounter := make(map[string]int) //练习题5.2
	links := visit(nil, doc, true, nameCounter)
	for i := range links {
		fmt.Printf("visit link=%s\n", links[i])
	}
	fmt.Printf("namecounter=%v", nameCounter)
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

//练习题5.1
//visit appends to links each link found in n and returns the result.
func visit(links []string, n *html.Node, first bool, counter map[string]int) []string {
	if n == nil {
		return links
	}

	counter[n.Data] ++
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				links = append(links, a.Val)
			}
		}
	} else if n.Type == html.TextNode {
		fmt.Printf("text node = %v\n", n)
	} else {
		fmt.Printf("other node = %v\n", n)
	}

	if first {
		links = visit(links, n.FirstChild, false, counter)
	} else {
		links = visit(links, n.NextSibling, false, counter)
	}
	return links
}

//练习题 5.19
func panic1() (r int) {
	defer func() {
		p := recover()
		r = p.(int)
	}()

	panic(5)
}

//练习题 6.1
type IntSet struct {
	words []uint
}

// Has reports whether the set contains the non-negative value x.
func (s *IntSet) Has(x int) bool {
	byteLen := s.ByteLen()
	word, bit := x/byteLen, uint(x%byteLen)
	return word < len(s.words) && s.words[word]&(1<<bit) != 0
}

// Add adds the non-negative value x to the set.
func (s *IntSet) Add(x int) {
	byteLen := s.ByteLen()
	word, bit := x/byteLen, uint(x%byteLen)
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	s.words[word] |= 1 << bit
}

// UnionWith sets s to the union of s and t.
func (s *IntSet) UnionWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] |= tword
		} else {
			s.words = append(s.words, tword)
		}
	}
}

// String returns the set as a string of the form "{1 2 3}".
func (s *IntSet) String() string {
	byteLen := s.ByteLen()
	var buf bytes.Buffer
	buf.WriteByte('{')
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < byteLen; j++ {
			if word&(1<<uint(j)) != 0 {
				if buf.Len() > len("{") {
					buf.WriteByte('}')
				}
				fmt.Fprintf(&buf, "%d", byteLen*i+j)
			}
		}
	}
	buf.WriteByte('}')
	return buf.String()
}

// return the number of elements
func (s *IntSet) Len() int {
	var r int
	for i := range s.words {
		r += pkg.PopCount2_5(uint64(s.words[i]))
	}
	return r
}

// remove x from the set
func (s *IntSet) Remove(x int) {
	if !s.Has(x) {
		return
	}
	byteLen := s.ByteLen()
	word, bit := x/byteLen, uint(x%byteLen)
	s.words[word] = ^((^s.words[word]) | (1 << bit))
}

// remove all elements from the set
func (s *IntSet) Clear() {
	s.words = make([]uint, 0)
}

// return a copy of the set
func (s *IntSet) Copy() *IntSet {
	words := make([]uint, 0, len(s.words))
	for i := range s.words {
		words[i] = s.words[i]
	}
	return &IntSet{words: words}
}

//练习题 6.2
func (s *IntSet) AddAll(xs ...int) {
	for i := range xs {
		s.Add(xs[i])
	}
}

//练习题 6.3
func (s *IntSet) IntersectWith(t *IntSet) {
	len1 := len(s.words)
	len2 := len(t.words)
	for i := 0; i < len1 && i < len2; i++ {
		s.words[i] &= t.words[i]
	}
}

func (s *IntSet) DifferenceWith(t *IntSet) {
	len1 := len(s.words)
	len2 := len(t.words)
	if len1 > len2 {
		len1 = len2
	}
	for i := 0; i < len1; i++ {
		s.words[i] ^= s.words[i] & t.words[i]
	}
}

func (s *IntSet) SymmetricDifference(t *IntSet) {
	s1 := s.Copy()
	s.UnionWith(t)
	s1.IntersectWith(t)
	s.DifferenceWith(s1)
}

//练习题6.4
func (s *IntSet) Elems() []int {
	byteLen := s.ByteLen()
	r := make([]int, 0, s.Len())
	for i := range s.words {
		for j := 0; j < byteLen; j++ {
			if s.words[i]&(1<<uint(j)) != 0 {
				r = append(r, i*byteLen+j)
			}
		}
	}
	return r
}

//练习题6.5
func (s *IntSet) ByteLen() int {
	return 32 << (^uint(0) >> 63)
}

//练习题7.1
type ByteCounter struct {
	words int
	lines int
}

func (b *ByteCounter) String() string {
	return fmt.Sprintf("{words:%d,lines:%d}", b.words, b.lines)
}

func (b *ByteCounter) Write(p []byte) (int, error) {
	start := 0
	for {
		advance, _, _ := bufio.ScanWords(p[start:], true)
		if advance == 0 {
			break
		}

		start += advance
		b.words ++
	}

	start = 0
	for {
		advance, _, _ := bufio.ScanLines(p[start:], true)
		if advance == 0 {
			break
		}

		start += advance
		b.lines ++
	}
	return len(p), nil
}

//练习题7.2
type CW struct {
	count int64
	w     io.Writer
}

func (cw *CW) Write(p []byte) (int, error) {
	n, err := cw.w.Write(p)
	cw.count += int64(n)
	return n, err
}
func CountingWriter(w io.Writer) (io.Writer, *int64) {
	cw := &CW{}
	cw.w = w
	return cw, &cw.count
}
