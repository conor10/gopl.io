// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 254.
//!+

// Chat is a server that lets clients chat with each other.
package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"time"
)

//!+broadcaster
type client struct { // an outgoing message channel
	name string
	channel chan<- string
}

var (
	entering = make(chan client)
	leaving  = make(chan client)
	messages = make(chan string) // all incoming client messages
)

func broadcaster() {
	clients := make(map[client]bool) // all connected clients

	for {
		select {
		case msg := <-messages:
			// Broadcast incoming message to all
			// clients' outgoing message channels.
			for cli, _ := range clients {
				cli.channel <- msg
			}

		case cli := <-entering:
			cli.channel <- getClients(clients)
			clients[cli] = true

		case cli := <-leaving:
			delete(clients, cli)
			close(cli.channel)
		}
	}
}

func getClients(clients map[client]bool) string {
	clientList := make([]string, len(clients))
	for client := range clients {
		clientList = append(clientList, client.name)
	}

	return fmt.Sprintf("Active clients: [%s]", strings.Join(clientList, ","))
}

//!-broadcaster

//!+handleConn
func handleConn(conn net.Conn) {
	ch := make(chan string) // outgoing client messages
	go clientWriter(conn, ch)

	who := conn.RemoteAddr().String()
	id := client{who, ch}
	ch <- "You are " + who
	messages <- who + " has arrived"
	entering <- id

	msgSent := make(chan struct{})

	go func() {
		timeout := time.NewTicker(1 * time.Minute)

		loop:
			for {
				select {
				case <- timeout.C:
					conn.Close()
					break loop
				case <- msgSent:
					timeout.Stop()
					timeout = time.NewTicker(1 * time.Minute)
				}
			}
	}()

	input := bufio.NewScanner(conn)
	for input.Scan() {
		messages <- who + ": " + input.Text()
		msgSent <- struct{}{}
	}
	// NOTE: ignoring potential errors from input.Err()

	leaving <- id
	messages <- who + " has left"
	conn.Close()
}

func clientWriter(conn net.Conn, ch <-chan string) {
	for msg := range ch {
		fmt.Fprintln(conn, msg) // NOTE: ignoring network errors
	}
}

//!-handleConn

//!+main
func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}

	go broadcaster()
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)
	}
}

//!-main
