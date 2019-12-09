package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func newTestServer() *httptest.Server {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Write([]byte("ok"))
	})

	mux.HandleFunc("/missing", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(404)
	})

	return httptest.NewServer(mux)
}

func TestExecWithoutRealRequest(t *testing.T) {
	server := newTestServer()
	defer server.Close()

	testCases := []struct {
		url     string
		success bool
	}{
		{
			url:     server.URL,
			success: true,
		},
		{
			url:     server.URL + "/missing",
			success: false,
		},
		{
			url:     server.URL,
			success: true,
		},
	}

	var urls []string
	for _, tes := range testCases {
		urls = append(urls, tes.url)
	}

	requester := newRequester(urls)
	requester.Exec()

	for i, r := range requester.requests {
		tc := testCases[i]

		if tc.success {
			if r.err != nil {
				t.Errorf("%s : expect success, but fail", r.URL)
			}
		} else {
			if r.err == nil {
				t.Errorf("%s : expect fail but success", r.URL)
			}
		}
	}
}

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
