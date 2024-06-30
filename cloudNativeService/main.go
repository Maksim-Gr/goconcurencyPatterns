package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func helloSimpleHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello simple handler "))
}
func helloGorillaMuxHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from gorilla mux handler"))
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/gorilla", helloGorillaMuxHandler)

	http.HandleFunc("/", helloSimpleHandler)
	log.Fatal(http.ListenAndServe(":8080", r))
}
