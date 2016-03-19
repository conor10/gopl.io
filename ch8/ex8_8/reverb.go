// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 224.

// Reverb2 is a TCP server that simulates an echo.
package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"time"
)

func echo(c net.Conn, shout string, delay time.Duration) {
	fmt.Fprintln(c, "\t", strings.ToUpper(shout))
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", shout)
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", strings.ToLower(shout))
}

//!+
func handleConn(c net.Conn) {
	data := make(chan string)

	go func() {
		input := bufio.NewScanner(c)
		for input.Scan() {
			data <- input.Text()
		}
	}()

	for {
		select {
		case text := <- data:
			go echo(c, text, 1*time.Second)
		case <-time.After(10 * time.Second):
			fmt.Printf("Timeout disconnecting client: %q", c)
			close(data)
			c.Close()
			return
		}
	}



	// NOTE: ignoring potential errors from input.Err()
}

//!-

func main() {
	l, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := l.Accept()
		if err != nil {
			log.Print(err) // e.g., connection aborted
			continue
		}
		go handleConn(conn)
	}
}
