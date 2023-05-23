package main

import (
	"flag"
	"log"
	"net/http"
	"os"

	"goForum.oksuriini.net/internal/models"
)

type application struct {
	errorLogger *log.Logger
	infoLogger  *log.Logger
	messages    *models.MessageModel
	threads     *models.ThreadModel
	subjects    *models.SubjectModel
}

func main() {

	// address for port number and dsn for you database dsn
	addr := flag.String("addr", ":4050", "Port number from which the application servers")
	dsn := flag.String("dsn", "web:salis@/goforum?parseTime=true", "Database DSN")
	flag.Parse()

	// create loggers for loggin info and errors
	infoLogger := log.New(os.Stdout, "INFO \t", log.Ldate|log.Ltime)
	errorLogger := log.New(os.Stdout, "ERROR \t", log.Llongfile|log.Ldate|log.Ltime)

	// openDB func in dbfunc.go
	db, err := openDB(*dsn, "mysql")
	defer db.Close()

	// application struct holds
	app := &application{
		errorLogger: errorLogger,
		infoLogger:  infoLogger,
		messages:    &models.MessageModel{DB: db},
	}

	//srv to hold Server struct
	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: errorLogger,
		Handler:  app.routes(),
	}

	// inform about server getting started on number 'addr'
	infoLogger.Printf("Starting server on %s\n", *addr)
	err = srv.ListenAndServe()
	if err != nil {
		errorLogger.Fatal(err)
	}
}
