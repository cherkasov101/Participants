package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/answers", answers)
	mux.HandleFunc("/participants", participants)

	log.Fatal(http.ListenAndServe(":8080", mux))
}
