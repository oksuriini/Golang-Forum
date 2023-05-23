package main

import "net/http"

func (app *application) routes() *http.ServeMux {

	mux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir("./ui/static/"))

	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	mux.HandleFunc("/", app.home)
	mux.HandleFunc("/forum", app.forum)

	// TODO LIST:
	// mux.HandleFunc("/forum/subject", viewSubject)
	// mux.HandleFunc("/forum/subject/thread", viewThread)
	// mux.HandleFunc("/forum/subject/thread", postThread)

	return mux
}
