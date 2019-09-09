package pkg

import (
	"math/rand"
	"testing"
	"time"
)

/**
文件必须以 *_test.go 命名
函数必须以Test* 命名并且带参数 *testing.T
*/

func charcount2(s string) int {
	var count int
	for range s {
		count++
	}
	return count
}

//练习题11.1 功能测试。
func TestCharCount(t *testing.T) {
	var tests = []struct {
		input string
		want  int
	}{
		{"我", 1},
		{"abc", 3},
	}

	for i := range tests {
		if v := charcount2(tests[i].input); v != tests[i].want {
			//如果想要程序遇到错误就中止可以使用t.Fatalf
			t.Errorf("input(%s)=%d, want %d\n", tests[i].input, v, tests[i].want)
		}
	}
}

//可以编写多个Test，每个都会被执行
func TestIntSet(t *testing.T) {

}

func randomPalindrome2(rng *rand.Rand) string {
	n := rng.Intn(25) // random length up to 24
	runes := make([]rune, n)
	for i := 0; i < (n+1)/2; i++ {
		r := rune(rng.Intn(0x1000)) // random rune up to '\u0999'
		runes[i] = r
		runes[n-1-i] = r + rune(rng.Intn(0x7fff))
	}
	runes2 := make([]rune, n+2)
	runes2 = append(runes2, 'a')
	copy(runes2[1:], runes)
	runes2 = append(runes2, 'b')
	return string(runes2)
}

//练习题11.3
func TestRandomPalindromes(t *testing.T) {
	seed := time.Now().UTC().UnixNano()
	t.Logf("Random seed: %d", seed)
	rng := rand.New(rand.NewSource(seed))
	for i := 0; i < 1000; i++ {
		p := randomPalindrome2(rng)
		if IsPalindrome(p) {
			t.Errorf("IsPalindrome(%q) = true", p)
		}
	}
}

/**
基准测试和普通测试差不多，但是以Benchmark为前缀名
 -test.bench=. 等价于 直接在goland编辑器里面 TestFramework中选择gobench

-test.bench=RandomPalindromes
-test.benchmem
*/
func BenchmarkRandomPalindromes(t *testing.B) {
	/**
	BenchmarkRandomPalindromes-4   	 2000000	       619 ns/op
	结果中基准测试名的数字后缀部分 BenchmarkRandomPalindromes-4，这里是4，表示运行时对应的GOMAXPROCS的值
	后面表示执行2000000次平均时间是 619纳秒/次

	增加命令参数 -test.benchmem
	BenchmarkRandomPalindromes-4   	 3000000	       405 ns/op	     128 B/op	       1 allocs/op
	*/
	for i := 0; i < t.N; i++ {
		IsPalindrome("A man, a plan, a canal: Panama")
	}
}
