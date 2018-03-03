package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	db := database{"shoes": 50, "socks": 20}
	log.Fatal(http.ListenAndServe("localhost:2000", db))
}

type dollars float32

func (d dollars) String() string { return fmt.Sprintf("$%2.f", d) }

type database map[string]int

// db now implements http.Handler interface
func (db database) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	switch req.URL.Path {
	case "/list":
		for item, price := range db {
			fmt.Fprintf(w, "%s: %d\n", item, price)
		}
	case "/price":
		item := req.URL.Query().Get("item")
		price, ok := db[item]
		if !ok {
			// w.WriteHeader(http.StatusNotFound) // 404
			msg := fmt.Sprintf("no such item: %q\n", item)
			http.Error(w, msg, http.StatusNotFound)
			return
		}
		fmt.Fprintf(w, "%d\n", price)
	default:
		// w.WriteHeader(http.StatusNotFound) // 404
		msg := fmt.Sprintf("no such page: %s\n", req.URL)
		http.Error(w, msg, http.StatusNotFound)
	}
}
