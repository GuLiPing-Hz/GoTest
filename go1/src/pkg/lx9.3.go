//实现一个缓冲无阻塞的服务器
package pkg

import (
	"fmt"
	"pkg/B"
	"pkg/internal/A"
	"sync"
)

type Func func(key string) (interface{}, error)

type result struct {
	value interface{}
	err   error
}

type entry struct {
	res   result
	ready chan struct{} // closed when res is ready
}

type Memo struct {
	f     Func
	mu    sync.Mutex // guards cache
	cache map[string]*entry
}

func New(f Func) *Memo {
	fmt.Printf("A=%d\n", A.A)
	//fmt.Printf("a=%d\n", A.a) //小写没有导出，无法访问。
	fmt.Printf("B=%d\n", B.B)
	//fmt.Printf("b=%d\n", B.b) //小写没有导出，无法访问。
	return &Memo{f: f, cache: make(map[string]*entry)}
}

func (memo *Memo) Get(key string) (interface{}, error) {
	memo.mu.Lock()
	e, ok := memo.cache[key]
	if !ok {
		e = &entry{result{}, make(chan struct{})}
		memo.cache[key] = e
		memo.mu.Unlock()

		e.res.value, e.res.err = memo.f(key)
		close(e.ready)
		return e.res.value, e.res.err
	}
	memo.mu.Unlock()

	<-e.ready
	return e.res.value, e.res.err
}

type Memo2 struct {
	f       Func
	cache   map[string]*entry
	request chan Request
}

type Request struct {
	key      string
	response chan result
	done     chan struct{}
}

func New2(f Func) *Memo2 {
	ret := &Memo2{f: f, cache: make(map[string]*entry), request: make(chan Request),}
	go ret.background()
	return ret
}

func (memo *Memo2) Get2(key string, done chan struct{}) (interface{}, error) {
	resp := make(chan result)
	memo.request <- Request{key, resp, done}
	r := <-resp
	return r.value, r.err
}

func (memo *Memo2) background() {
	for req := range memo.request {
		if e, ok := memo.cache[req.key]; ok {
			go func() {
				select {
				case <-e.ready:
					//这里等待数据加载完毕，信道关闭表示已经加载完毕了
				case <-req.done:
					e.res.value = nil
					e.res.err = fmt.Errorf("canceled")
				}

				req.response <- e.res
			}()
			continue
		}

		e := entry{ready: make(chan struct{})}
		memo.cache[req.key] = &e

		go func() {
			signal := make(chan result)
			go func() {
				res := result{}
				res.value, res.err = memo.f(req.key)
				signal <- res
			}()

			select {
			case e.res = <-signal:
			case <-req.done:
				e.res.value = nil
				e.res.err = fmt.Errorf("canceled")
			}

			close(e.ready)
			req.response <- e.res
		}()
	}
}
