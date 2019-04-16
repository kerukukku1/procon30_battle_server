package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func entryHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>entry point</h1>"))
}

func main() {
	// http.HandleFunc("/entry", entryHandler)
	// http.ListenAndServe(":8080", nil)
	router := mux.NewRouter()
	router.HandleFunc("/entry", entryHandler).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", router))
}
