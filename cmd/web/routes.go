package main

import "net/http"

func (app *application) routes() *http.ServeMux {

	mux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir("./ui/static/"))

	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	mux.HandleFunc("/", app.home)
	mux.HandleFunc("/forum", app.forum)
	mux.HandleFunc("/forum/create", app.createMessage)
	mux.HandleFunc("/forum/subjects", app.getSubjects)
	mux.HandleFunc("/forum/subject", app.getThreads)
	mux.HandleFunc("/forum/thread", app.getThreadMessages)

	// TODO LIST:
	//
	// mux.HandleFunc("/forum/subject/thread", viewThread)
	// mux.HandleFunc("/forum/subject/thread", postThread)

	return mux
}
