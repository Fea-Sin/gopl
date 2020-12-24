/*
/*
@Time : 2020/12/24 3:12 下午
@Author : chengqunzhong
@File : memo
@Software: GoLand
*/
package memo2

import "sync"

type Func func(key string) (interface{}, error)

type result struct {
	value interface{}
	err error
}

type entry struct {
	res result
	ready chan struct{}
}

type Memo struct {
	f Func
	mu sync.Mutex
	cache map[string]*entry
}

func (memo *Memo) Get(key string) (value interface{}, err error) {
	memo.mu.Lock()
	e := memo.cache[key]
	if e == nil {
		e = &entry{ready: make(chan struct{})}
		memo.cache[key] = e
		memo.mu.Unlock()

		e.res.value, e.res.err = memo.f(key)
		close(e.ready)
	} else {
		memo.mu.Unlock()
		<-e.ready
	}
	return e.res.value, e.res.err
}

func New(f Func) *Memo {
	return &Memo{f: f, cache: make(map[string]*entry)}
}