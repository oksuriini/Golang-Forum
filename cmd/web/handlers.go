package main

import (
	"html/template"
	"log"
	"net/http"
)

// Add function handlers here

func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	files := []string{
		"./ui/html/base.tmpl.html",
		"./ui/html/partials/footer.tmpl.html",
		"./ui/html/partials/nav.tmpl.html",
		"./ui/html/pages/home.tmpl.html",
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	err = ts.ExecuteTemplate(w, "base", nil)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
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
//
//func viewSubject(w http.ResponseWriter, r *http.Request) {
//	subject := r.URL.Query().Get("sub")
//	if subject == "" {
//		http.NotFound(w, r)
//		return
//	}
//	fmt.Fprint(w, fmt.Sprintf("Displaying subject: %s", subject))
//}

// func viewSubfolder(w http.ResponseWriter, r *http.Request) {}
//
// func postSubfolder(w http.ResponseWriter, r *http.Request) {}

//func viewThread(w http.ResponseWriter, r *http.Request) {
//	if r.Method != http.MethodGet {
//		w.Header().Set("Allow", http.MethodGet)
//		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
//	}
//	w.Write([]byte("Here are thread messages"))
//}
//
//func postThread(w http.ResponseWriter, r *http.Request) {
//	if r.Method != http.MethodPost {
//		w.Header().Set("Allow", http.MethodPost)
//		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
//	}
//	w.Write([]byte("Message posted successfully"))
//}
