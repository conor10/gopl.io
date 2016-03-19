// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 195.

// Http4 is an e-commerce server that registers the /list and /price
// endpoint by calling http.HandleFunc.
package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

//!+main

func main() {
	db := database{"shoes": 50, "socks": 5}
	http.HandleFunc("/list", db.list)
	http.HandleFunc("/price", db.price)
	http.HandleFunc("/create", db.create)
	http.HandleFunc("/update", db.update)
	http.HandleFunc("/delete", db.delete)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

//!-main

type dollars float32

func (d dollars) String() string { return fmt.Sprintf("$%.2f", d) }

type database map[string]dollars

func (db database) list(w http.ResponseWriter, req *http.Request) {
	for item, price := range db {
		fmt.Fprintf(w, "%s: %s\n", item, price)
	}
}

func (db database) price(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	if price, ok := db[item]; ok {
		fmt.Fprintf(w, "%s\n", price)
	} else {
		w.WriteHeader(http.StatusNotFound) // 404
		fmt.Fprintf(w, "no such item: %q\n", item)
	}
}

func (db database) create(w http.ResponseWriter, req *http.Request) {
	params := req.URL.Query()
	item := params.Get("item")
	price := params.Get("price")
	f, err := strconv.ParseFloat(price, 32)
	if err == nil {
		priceDollars := dollars(f)
		db[item] = priceDollars
		fmt.Fprintf(w, "created value %s:%s", item, priceDollars)
	} else {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "invalid price: %s\n", price)
	}
}

func (db database) update(w http.ResponseWriter, req *http.Request) {
	params := req.URL.Query()
	item := params.Get("item")
	price := params.Get("price")
	f, err := strconv.ParseFloat(price, 32)
	if err == nil {
		priceDollars := dollars(f)
		db[item] = priceDollars
		fmt.Fprintf(w, "updated value %s:%s", item, priceDollars)
	} else {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "invalid price: %s\n", price)
	}
}

func (db database) delete(w http.ResponseWriter, req *http.Request) {
	params := req.URL.Query()
	item := params.Get("item")
	if _, ok := db[item]; ok {
		delete(db, item)
	}
	fmt.Fprintf(w, "deleted item %s\n", item)
}
