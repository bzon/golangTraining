package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

type database map[string]int

func main() {
	db := database{"shoes": 50, "rolex": 1000}

	// it’s typical to define HTTP handlers across many files of an application, and it would be a nuisance if they all had to be explicitly registered with the application’s ServeMux instance.
	// So, for convenience, net/http provides a global ServeMux instance called DefaultServeMux and package-level functions called http.Handle and http.HandleFunc. To use DefaultServeMux as the server’s main handler, we needn’t pass it to ListenAndServe; nil will do.

	// Donovan, Alan A. A.. The Go Programming Language (Addison-Wesley Professional Computing Series) (p. 195). Pearson Education. Kindle Edition.
	http.HandleFunc("/home", home)
	http.HandleFunc("/list", db.list)
	http.HandleFunc("/price", db.price)
	http.HandleFunc("/update", db.update)
	http.HandleFunc("/create", db.create)
	http.HandleFunc("/delete", db.delete)
	log.Fatal(http.ListenAndServe("localhost:2000", nil)) // Use nil
}

func (db database) price(w http.ResponseWriter, r *http.Request) {
	item := r.URL.Query().Get("item")
	price, ok := db[item]
	if !ok {
		msg := fmt.Sprintf("no such item: %q\n", item)
		http.Error(w, msg, http.StatusNotFound)
		return
	}
	fmt.Fprintf(w, "%d\n", price)
}

func (db database) list(w http.ResponseWriter, r *http.Request) {
	for item, price := range db {
		fmt.Fprintf(w, "%s: %d\n", item, price)
	}
}

// Exercise 7.11: Add additional handlers so that clients can create, read, update, and delete database entries. For example, a request of the form /update?item=socks&price=6 will update the price of an item in the inventory and report an error if the item does not exist or if the price is invalid. (Warning: this change introduces concurrent variable updates.)
func (db database) update(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	item := params.Get("item")
	_, ok := db[item]
	if !ok {
		msg := fmt.Sprintf("no such item %s", item)
		http.Error(w, msg, http.StatusBadRequest)
		return
	}
	newPrice, err := strconv.Atoi(params.Get("price"))
	if err != nil {
		http.Error(w, "Failed", http.StatusExpectationFailed)
		return
	}
	log.Printf("%s Price updated to %d\n", item, newPrice)
	db[item] = newPrice
}

func (db database) create(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	item := params.Get("item")
	_, ok := db[item]
	if ok {
		msg := fmt.Sprintf("Item %s already exists!", item)
		http.Error(w, msg, http.StatusBadRequest)
		return
	}
	price := params.Get("price")
	intPrice, err := strconv.Atoi(price)
	if intPrice == 0 {
		http.Error(w, "Missing price!", http.StatusExpectationFailed)
		return
	}
	if err != nil {
		http.Error(w, "Failed", http.StatusExpectationFailed)
		return
	}
	db[item] = intPrice
	fmt.Fprintf(w, "%s added to database\n", item)
	log.Printf("%s added to database\n", item)
}

func (db database) delete(w http.ResponseWriter, r *http.Request) {
	item := r.URL.Query().Get("item")
	_, ok := db[item]
	if !ok {
		msg := fmt.Sprintf("The item %s does not exists", item)
		http.Error(w, msg, http.StatusBadRequest)
		return
	}
	delete(db, item)
	fmt.Fprintf(w, "%s deleted from the database!\n", item)
	log.Printf("%s deleted from the database!\n", item)
}

func home(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "welcome!\n")
}
