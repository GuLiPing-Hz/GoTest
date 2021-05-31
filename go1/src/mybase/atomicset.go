package mybase

import (
	"math/rand"
	"sync"
	"sync/atomic"
)

type AtomicSet struct {
	set sync.Map
	len int32
}

func (imp *AtomicSet) Range(cb func(val interface{}) bool) {
	if cb == nil {
		return
	}

	imp.set.Range(func(key, value interface{}) bool {
		return cb(key)
	})
}

func (imp *AtomicSet) Insert(val interface{}) {
	hasIt := imp.Contain(val)
	imp.set.Store(val, true)
	if !hasIt {
		atomic.AddInt32(&imp.len, 1)
	}
}

func (imp *AtomicSet) Remove(val interface{}) {
	if imp.Contain(val) {
		atomic.AddInt32(&imp.len, -1)
	}
	imp.set.Delete(val)
}

func (imp *AtomicSet) Contain(val interface{}) bool {
	_, ok := imp.set.Load(val)
	return ok
}

func (imp *AtomicSet) Len() int {
	return int(atomic.LoadInt32(&imp.len))
}

func (imp *AtomicSet) Random() (interface{}, bool) {
	var idleLen = int32(imp.Len())
	var index int32
	var ret interface{}

	if idleLen > 0 {
		r := rand.Int31n(idleLen)

		index = 0
		imp.Range(func(val interface{}) bool {
			if index == r {
				ret = val
				return false
			}

			index++
			return true
		})
		return ret, true
	} else {
		return nil, false
	}
}
