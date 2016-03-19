// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

package memo_test

import (
	"fmt"
	"log"
	"sync"
	"time"
	"testing"

	"gopl.io/ch9/ex9_3"
	"gopl.io/ch9/memotest"
)

var httpGetBody = memotest.HTTPGetBody

func Test(t *testing.T) {
	cancel := make(chan struct{})
	m := memo.New(memo.Func{Cancel: cancel, F: httpGetBody})
	defer m.Close()
	memotest.Sequential(t, m)
}

func TestConcurrent(t *testing.T) {
	cancel := make(chan struct{})
	m := memo.New(memo.Func{Cancel: cancel, F: httpGetBody})
	defer m.Close()
	memotest.Concurrent(t, m)
}

func TestCancel(t *testing.T) {
	cancel := make(chan struct{})
	m := memo.New(memo.Func{Cancel: cancel, F: httpGetBody})
	defer m.Close()

	var n sync.WaitGroup
		n.Add(1)
		go func(url string) {
			start := time.Now()
			_, err := m.Get(url)
			if err != nil {
				log.Print(err)
			}
			fmt.Printf("%s, %s\n",
				url, time.Since(start))
			n.Done()
		}("https://golang.org")
		go func() {
			cancel <- struct{}{}
		}()
	n.Wait()
}