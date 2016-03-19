package main

import (
	"bufio"
	"io"
	"os"
	"fmt"
)

func main() {
	text := "Some example text"
	w := bufio.NewWriter(os.Stdout)
	cw, written := CreateCountingWriter(w)
	fmt.Println(*written)
	cw.Write([]byte(text))
	fmt.Println(*written)
}

type CountingWriter struct {
	io.Writer
	int64
}

func (c *CountingWriter) Write(p []byte) (n int, err error) {
	n, err = c.Writer.Write(p)
	c.int64 += int64(n)
	return n, err
}

func CreateCountingWriter(w io.Writer) (io.Writer, *int64) {
	cw := CountingWriter{w, 0}
	return &cw, &cw.int64
}
