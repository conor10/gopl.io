package main

import (
	"fmt"
	"log"
	"io"
	"net"
	"os"
	"time"
	"strings"
	"bufio"
)

func main() {
	for _, arg := range os.Args[1:] {
		details := strings.Split(arg, "=")
		if len(details) == 2 {
			go run(details[0], details[1])
		}
	}
	for {
		time.Sleep(60 * time.Second)
	}
}

func run(location, host string) {
	conn, err := net.Dial("tcp", host)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	prefix := fmt.Sprintf("[%s]: ", location)

	input := bufio.NewScanner(conn)
	for input.Scan() {
		log.Println(prefix + input.Text())
	}
}