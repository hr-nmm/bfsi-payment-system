package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	log.Println("server running, logs appear here...")
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		d, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(rw, "oops!!", http.StatusBadRequest)
			return
		}
		fmt.Fprintf(rw, "passed data => %s\n", d)
	})
	http.HandleFunc("/goodbye", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Goodbye world")
	})
	http.ListenAndServe("localhost:4221", nil)
}
