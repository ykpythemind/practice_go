package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
)

type request struct {
	URL          string
	responseBody string
	err          error
}

type requester struct {
	requests []*request
}

func (r *request) get() {
	// time.Sleep(1 * time.Second)

	// t := time.Now()
	// // store result
	// r.responseBody = t.Format("Mon Jan 2 15:04:05 -0700 MST 2006")

	resp, err := http.Get(r.URL)
	if err != nil {
		r.err = err
		return
	}
	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		r.err = err
		return
	}
	r.responseBody = string(b)
}

func newRequester(targetURL []string) requester {
	var req []*request
	for _, v := range targetURL {
		r := request{URL: v}
		req = append(req, &r)
	}

	return requester{requests: req}
}

func (r *requester) Exec() {
	ch := make(chan bool, 5) // 同時リスエスト数
	defer close(ch)

	var wg sync.WaitGroup

	for i, r := range r.requests {
		ch <- true // バッファがいっぱいになっていたらここでブロックしてくれる. 5つ実行中でいっぱいになる

		wg.Add(1)
		// execute request
		go func(r *request, i int) {
			fmt.Printf("try request / n=%d\n", i)
			r.get()

			wg.Done()
			<-ch // おわったらここで読み出してやる. チャンネルに空きができ、ループのブロックが解除される
		}(r, i)
	}

	wg.Wait() // すべてのリクエストが実行されるまで待つ
}
