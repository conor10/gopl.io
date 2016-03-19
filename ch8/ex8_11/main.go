package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
)

var done = make(chan struct{})

func cancelled() bool {
	select {
	case <-done:
		return true
	default:
		return false
	}
}


//!+
// Fetch downloads the URL and returns the
// name and length of the local file.
func process(resp http.Response) (filename string, n int64, err error) {
	local := path.Base(resp.Request.URL.Path)
	if local == "/" {
		local = "index.html"
	}
	f, err := os.Create(local)
	if err != nil {
		return "", 0, err
	}
	n, err = io.Copy(f, resp.Body)
	// Close file, but prefer error from Copy, if any.
	if closeErr := f.Close(); err == nil {
		err = closeErr
	}
	return local, n, err
}

//!-

type resource struct {
	url string
	resp *http.Response
}

func main() {
	responses := make(chan *resource)
	for _, url := range os.Args[1:] {
		go func(url string) {
			resp, err := http.Get(url)
			if err != nil {
				return
			}
			if !cancelled() {
				responses <- &resource{url, resp}
			} else {
				return
			}
		}(url)
	}

	select {
	case resource := <-responses:
		local, n, err := process(*resource.resp)
		resource.resp.Body.Close()
		close(done)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch %s: %v\n", resource.url, err)
		}
		fmt.Fprintf(os.Stderr, "%s => %s (%d bytes).\n", resource.url, local, n)
	case <-done:
		for r := range responses {
			r.resp.Body.Close()
		}
	}
}
