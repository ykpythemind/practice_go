package main

import (
	"fmt"
	"testing"
)

func TestExec(t *testing.T) {
	urls := []string{"https://twitter.com/ykpythemind", "https://ykpythemind.com", "https://instagram.com", "https://example.com", "invalidaddress", "https://golang.org/pkg/sync/#WaitGroup", "https://github.com/", "https://its_not_found_hogehogehogehoge.com"}
	requester := newRequester(urls)

	requester.Exec()

	for _, r := range requester.requests {
		if r.err == nil {
			var str string
			if len(r.responseBody) > 100 {
				str = r.responseBody[:100]
			} else {
				str = r.responseBody
			}
			fmt.Printf("success (%s) %s\n", r.URL, str)
		} else {
			fmt.Printf("fail %s\n", r.err)
		}
	}
}
