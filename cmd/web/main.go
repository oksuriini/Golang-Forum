package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/forum", forum)
	// mux.HandleFunc("/forum/subject", viewSubject)
	// mux.HandleFunc("/forum/subject/thread", viewThread)
	// mux.HandleFunc("/forum/subject/thread", postThread)

	log.Println("Starting Server on :4050")
	err := http.ListenAndServe(":4050", mux)
	if err != nil {
		log.Fatal(err)
	}
}
