package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

type application struct {
	errorLogger *log.Logger
	infoLogger  *log.Logger
}

func main() {

	addr := flag.String("addr", ":4050", "Port number from which the application servers")
	flag.Parse()

	infoLogger := log.New(os.Stdout, "INFO \t", log.Ldate|log.Ltime)
	errorLogger := log.New(os.Stdout, "ERROR \t", log.Llongfile|log.Ldate|log.Ltime)

	app := &application{
		errorLogger: errorLogger,
		infoLogger:  infoLogger,
	}

	mux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir("./ui/static/"))

	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	mux.HandleFunc("/", app.home)
	mux.HandleFunc("/forum", app.forum)

	// TODO LIST:
	// mux.HandleFunc("/forum/subject", viewSubject)
	// mux.HandleFunc("/forum/subject/thread", viewThread)
	// mux.HandleFunc("/forum/subject/thread", postThread)

	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: errorLogger,
		Handler:  mux,
	}

	infoLogger.Printf("Starting server on %s\n", *addr)
	err := srv.ListenAndServe()
	if err != nil {
		errorLogger.Fatal(err)
	}
}
