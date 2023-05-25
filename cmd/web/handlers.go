package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type Data struct {
	CurrentYear int
}

// Add function handlers here

func (app *application) home(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/" {
		app.notFound(w)
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
		app.serverError(w, err)
		return
	}

	err = ts.ExecuteTemplate(w, "base", nil)
	if err != nil {
		app.serverError(w, err)
	}
}

func (app *application) forum(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/forum" {
		app.notFound(w)
		return
	}

	files := []string{
		"./ui/html/base.tmpl.html",
		"./ui/html/partials/footer.tmpl.html",
		"./ui/html/partials/nav.tmpl.html",
		"./ui/html/pages/forum.tmpl.html",
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		app.serverError(w, err)
		return
	}

	err = ts.ExecuteTemplate(w, "base", nil)
	if err != nil {
		app.serverError(w, err)
	}
}

func (app *application) createMessage(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		app.clientError(w, http.StatusMethodNotAllowed)
		return
	}

	content := "Test Content"
	threadId := 1
	creatorId := 2

	app.messages.Insert(threadId, content, creatorId)

	http.Redirect(w, r, fmt.Sprintf("/"), http.StatusSeeOther)
}

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

//func register(w http.ResponseWriter, r *http.Request) {}
//
//func registerPost(w http.ResponseWriter, r *http.Request) {}
//
//func login(w http.ResponseWriter, r *http.Request) {}
//
//func logout(w http.ResponseWriter, r *http.Request) {}
//
