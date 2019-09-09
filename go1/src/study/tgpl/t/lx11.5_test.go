//一个main包里面，只能有一个main函数，我为了简单的做练习题，
// 把所有main函数的文件都放在一个目录，导致main函数多次申明，所以本文件无法测试
package t

import (
	"strings"
	"testing"
)

//练习题11.5
func TestSplit(t *testing.T) {
	tests := []struct {
		input string
		sep   string
		want  int
	}{{"a,b,c", ",", 3}}

	for _, v := range tests {
		if val := len(strings.Split(v.input, v.sep)); val != v.want {
			t.Errorf("input(%s,%s)=%d,want %d", v.input, v.sep, val, v.want)
		}
	}
}
