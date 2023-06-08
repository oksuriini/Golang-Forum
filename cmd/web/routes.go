package main

import "net/http"

func (app *application) routes() *http.ServeMux {

	mux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir("./ui/static/"))

	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	// basic handlers
	mux.HandleFunc("/", app.home)
	mux.HandleFunc("/forum", app.forum)

	// get handlers
	mux.HandleFunc("/forum/subjects", app.getSubjects)
	mux.HandleFunc("/forum/subject", app.getThreads)
	mux.HandleFunc("/forum/thread", app.getThreadMessages)

	// create handlers
	mux.HandleFunc("/forum/createsubject", app.createSubject)
	mux.HandleFunc("/forum/createthread", app.createThread)
	mux.HandleFunc("/forum/create", app.createMessage)

	// user handlers
	mux.HandleFunc("/forum/registrar", app.registerUser)
	mux.HandleFunc("/forum/register", app.registerUserPost)
	mux.HandleFunc("/forum/login", app.loginUser)
	mux.HandleFunc("/forum/loginpost", app.loginUserPost)

	// TODO
	// login
	// additional stuff??
	//

	//INFUTURE
	// admin page
	// docker integration/format etc -> whole program to work in docker out of the box

	return mux
}
