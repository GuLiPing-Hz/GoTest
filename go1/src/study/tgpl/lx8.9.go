package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"sync"
	"time"
)

type GB2 int64

func (m GB2) String() string {
	return fmt.Sprintf("%0.2f GB", float64(m)/1024/1024)
}

type DirSize struct {
	dir  string
	size GB2
}

var Token = make(chan bool, 20)
var WG sync.WaitGroup

func walkDir(base, dir string, report chan<- DirSize) {
	WG.Add(1)
	Token <- true //同时最多运行20个goroutine
	defer func() {
		<-Token
		WG.Done()
	}()

	infos, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Printf("err=%s\n", err.Error())
		return
	}

	for i := range infos {
		data := infos[i]
		if data.IsDir() {
			//fmt.Printf("info=%v\n", data.Name())
			subdir := filepath.Join(dir, data.Name())
			base2 := base
			if base2 == "" {
				base2 = subdir
			}
			go walkDir(base2, subdir, report)
		} else {
			base2 := base
			if base2 == "" {
				base2 = dir
			}

			WG.Add(1)
			go func() {
				defer WG.Done()
				report <- DirSize{base2, GB2(data.Size())}
			}()
		}
	}
}

func printDir(dict map[string]DirSize) {
	for k, v := range dict {
		fmt.Printf("dir=%s size=%s\n", k, v.size)
	}
	//下面这种方法size无法打印出 xx GB,只是个普通的int64数字，可以对 map 类型重定义，并实现Stringer的String方法
	//fmt.Printf("dirs = %v\n", dict)
}

func lx8_9() {
	report := make(chan DirSize)
	walkDir("", "D:\\", report)

	go func() {
		WG.Wait()
		//由于我们这里的运算比较慢，所以采用等待的方式
		for len(report) > 0 {
			time.Sleep(time.Second)
		}
		close(report)
	}()

	chanClock := time.NewTicker(time.Second)
	dirMap := make(map[string]DirSize)

	for {
		select {
		case <-chanClock.C:
			printDir(dirMap)
		case data, ok := <-report:
			if !ok {
				goto finish
			}
			dict, ok := dirMap[data.dir]
			if ok {
				dirMap[data.dir] = DirSize{data.dir, dict.size + data.size}
			} else {
				dirMap[data.dir] = data
			}
		}
	}

finish:
	chanClock.Stop()
	fmt.Println("finish")
	printDir(dirMap)
}

type CancelHttpCtx struct {
	done chan struct{}
	err  error
}

func (imp *CancelHttpCtx) Deadline() (deadline time.Time, ok bool) {
	return time.Now().Add(time.Second * 30), true
}

func (imp *CancelHttpCtx) Done() <-chan struct{} {
	return imp.done
}

// If Done is not yet closed, Err returns nil.
// If Done is closed, Err returns a non-nil error explaining why:
// Canceled if the context was canceled
// or DeadlineExceeded if the context's deadline passed.
// After Err returns a non-nil error, successive calls to Err return the same error.
func (imp *CancelHttpCtx) Err() error {
	return imp.err
}

func (imp *CancelHttpCtx) Value(key interface{}) interface{} {
	return nil
}

func (imp *CancelHttpCtx) cancel() {
	imp.err = fmt.Errorf("canceled happen")
	imp.done <- struct{}{}
}

//用来判断，当前线程是否需要结束
func cancled(chanStop <-chan bool) bool {
	select {
	case _, _ = <-chanStop:
		return true
	default:
		return false
	}
}

func lx8_10(url string) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Printf("err=%s\n", err.Error())
		return
	}

	ctx := new(CancelHttpCtx)
	ctx.done = make(chan struct{})
	defer close(ctx.done)
	req = req.WithContext(ctx)

	go func() {
		os.Stdin.Read(make([]byte, 1))
		fmt.Println("finish by cmd")
		ctx.cancel()
		//取消这次请求
		fmt.Println("finish the request")
	}()
	response, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Printf("err=%s\n", err.Error())
		return
	}

	buf, err := ioutil.ReadAll(response.Body)
	response.Body.Close()
	if err != nil {
		fmt.Printf("err=%s\n", err.Error())
		return
	}
	fmt.Printf("body=%s\n", string(buf))
}

func main() {
	//lx8_9()
	lx8_10("http://www.android.com")
}
