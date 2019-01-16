package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World")
}

func about(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "About Page")
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", index)
	r.HandleFunc("/about", about)
	fmt.Println("Server starting...")
	http.ListenAndServe(":3000", r)
}
