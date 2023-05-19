package main

import "net/http"

func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	w.Write([]byte("Here is homepage"))
}

func forum(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Here is forumpage"))
}

//func register(w http.ResponseWriter, r *http.Request) {}
//
//func registerPost(w http.ResponseWriter, r *http.Request) {}
//
//func login(w http.ResponseWriter, r *http.Request) {}
//
//func logout(w http.ResponseWriter, r *http.Request) {}

func viewSubject(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("Here is subject folder"))
}

// func viewSubfolder(w http.ResponseWriter, r *http.Request) {}
//
// func postSubfolder(w http.ResponseWriter, r *http.Request) {}

func viewThread(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("Method not allowed"))
	}
	w.Write([]byte("Here are thread messages"))
}

func postThread(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("Method not allowed"))
	}
	w.Write([]byte("Message posted successfully"))
}
