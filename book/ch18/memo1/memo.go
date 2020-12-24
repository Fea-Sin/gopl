/*
/*
@Time : 2020/12/24 2:44 下午
@Author : chengqunzhong
@File : memo
@Software: GoLand
*/
package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type Memo struct {
	f Func
	cache map[string]result
}

type Func func(key string) (interface{}, error)

type result struct {
	value interface{}
	err error
}

func (memo *Memo) Get(key string) (interface{}, error) {
	res, ok := memo.cache[key]
	if !ok {
		res.value, res.err = memo.f(key)
		memo.cache[key] = res
	}
	return res.value, res.err
}

func New(f Func) *Memo {
	return &Memo{f: f, cache: make(map[string]result)}
}

// http 请求函数
func httpGetBody(url string) (interface{}, error) {
	resp , err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}

func main() {
	m := New(httpGetBody)

	// 请求url
	incomingURLs := []string {
		"https://golang.org",
		"https://godoc.org",
		"https://play.golang.org",
		"http://gopl.io",
	}

	for _, url := range incomingURLs {
		start := time.Now()
		value, err := m.Get(url)
		if err != nil {
			log.Print(err)
		}
		fmt.Printf("%s, %s, %d\n",
				url, time.Since(start), len(value.([]byte)))
	}
}
