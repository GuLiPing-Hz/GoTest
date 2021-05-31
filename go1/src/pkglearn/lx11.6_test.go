package pkglearn

import (
	"fmt"
	"testing"
)

/**
所有test命令参数
  -test.bench regexp
    	run only benchmarks matching regexp
  -test.benchmem
    	print memory allocations for benchmarks
  -test.benchtime d
    	run each benchmark for duration d (default 1s)

查看线程阻塞分析文件
  -test.blockprofile file
    	write a goroutine blocking profile to file
  -test.blockprofilerate rate
    	set blocking profile rate (see runtime.SetBlockProfileRate) (default 1)
  -test.count n
    	run tests and benchmarks n times (default 1)
  -test.coverprofile file
    	write a coverage profile to file
  -test.cpu list
    	comma-separated list of cpu counts to run each test with

查看cpu分析文件
  -test.cpuprofile file
    	write a cpu profile to file
  -test.failfast
    	do not start new tests after the first test failure
  -test.list regexp
    	list tests, examples, and benchmarks matching regexp then exit

查看内存分析文件
  -test.memprofile file
    	write an allocation profile to file
  -test.memprofilerate rate
    	set memory allocation profiling rate (see runtime.MemProfileRate)
  -test.mutexprofile string
    	write a mutex contention profile to the named file after execution
  -test.mutexprofilefraction int
    	if >= 0, calls runtime.SetMutexProfileFraction() (default 1)
  -test.outputdir dir
    	write profiles to dir
  -test.parallel n
    	run at most n tests in parallel (default 4)
  -test.run regexp
    	run only tests and examples matching regexp
  -test.short
    	run smaller test suite to save time
  -test.testlogfile file
    	write test action log to file (for use only by cmd/go)
  -test.timeout d
    	panic test binary after duration d (default 0, timeout disabled)
  -test.trace file
    	write an execution trace to file
  -test.v
    	verbose: print additional output


对于上面生成的profile(二进制文件)我们需要用下面的命令才能看
go tool pprof -text -nodecount=10 cpu.txt
go tool pprof -text -nodecount=10 mem.txt
go tool pprof -text -nodecount=10 block.txt

参数 -text 用于指定输出格式，在这里每行是一个函数，根据使用CPU的时间长短来排序。
其中 -nodecount=10 标志参数限制了只输出前10行的结果
*/

//-test.benchmem
func BenchmarkPopCount2_4(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCount2_4(326589914)
	}
}

//Test* Benchmark* Example*
func ExamplePopCount2_4() {
	//下面的注释代表了打印的结果，而且只能用 //注释 似乎只能这样。。
	fmt.Println(PopCount2_4(1))

	//Output:
	//1
}
