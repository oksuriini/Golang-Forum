package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"

	"goForum.oksuriini.net/internal/models"
)

type Data struct {
	CurrentYear int
}

type DataPost struct {
	Form createForm
}

type DataPass struct {
	Data        []*models.Message
	ThreadTitle string
}

type DataSubPass struct {
	Data         []*models.Thread
	SubjectID    int
	SubjectTitle string
}

type createForm struct {
	CreatorID   string `form:"creator"`
	Content     string `form:"content"`
	ThreadTitle string `form:"threadtitle"`
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

func (app *application) getThreadMessages(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/forum/thread" {
		app.notFound(w)
		return
	}

	thread := r.URL.Query().Get("thread")
	if r.Method != http.MethodGet {
		w.Header().Set("Allow", http.MethodGet)
		app.clientError(w, http.StatusMethodNotAllowed)
		return
	}

	threadTitle := r.URL.Query().Get("thread")

	if thread == "" {
		app.notFound(w)
		return
	}

	files := []string{
		"./ui/html/base.tmpl.html",
		"./ui/html/partials/footer.tmpl.html",
		"./ui/html/partials/nav.tmpl.html",
		"./ui/html/pages/messages.tmpl.html",
	}

	threadId, err := app.messages.GetThreadId(threadTitle)
	if err != nil {
		app.serverError(w, err)
		return
	}

	data, err := app.messages.GetMessagesInThread(threadId)
	if err != nil {
		app.serverError(w, err)
		return
	}

	finalData := DataPass{
		Data:        data,
		ThreadTitle: threadTitle,
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		app.serverError(w, err)
		return
	}

	err = ts.ExecuteTemplate(w, "base", finalData)
	if err != nil {
		app.serverError(w, err)
	}
}

func (app *application) getThreads(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/forum/subject" {
		app.notFound(w)
		return
	}

	subjectTitle := r.URL.Query().Get("subject")

	subjectId, err := app.messages.GetSubjectId(subjectTitle)
	if err != nil {
		app.serverError(w, err)
		return
	}

	files := []string{
		"./ui/html/base.tmpl.html",
		"./ui/html/partials/footer.tmpl.html",
		"./ui/html/partials/nav.tmpl.html",
		"./ui/html/pages/threads.tmpl.html",
	}

	data, err := app.messages.GetThreadsInSubject(subjectId)
	if err != nil {
		app.serverError(w, err)
		return
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		app.serverError(w, err)
		return
	}

	newData := DataSubPass{
		Data:         data,
		SubjectID:    subjectId,
		SubjectTitle: subjectTitle,
	}

	err = ts.ExecuteTemplate(w, "base", newData)
	if err != nil {
		app.serverError(w, err)
	}
}

func (app *application) getSubjects(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/forum/subjects" {
		app.notFound(w)
		return
	}

	files := []string{
		"./ui/html/base.tmpl.html",
		"./ui/html/partials/footer.tmpl.html",
		"./ui/html/partials/nav.tmpl.html",
		"./ui/html/pages/subjects.tmpl.html",
	}

	data, err := app.messages.GetAllSubjects()
	if err != nil {
		app.serverError(w, err)
		return
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		app.serverError(w, err)
		return
	}

	err = ts.ExecuteTemplate(w, "base", data)
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

	var form createForm

	err := r.ParseForm()
	if err != nil {
		app.serverError(w, err)
		return
	}

	app.formDecoder.Decode(form, r.Form)

	fmt.Println(r.FormValue("content"))

	content := r.FormValue("content")
	creatorID, err := strconv.Atoi(r.FormValue("creator"))
	if err != nil {
		app.serverError(w, err)
		return
	}
	tt := r.FormValue("threadtitle")

	fmt.Println(r.FormValue("threadtitle"))

	tid, err := app.messages.GetThreadId(r.FormValue("threadtitle"))
	if err != nil {
		app.serverError(w, err)
		return
	}

	fmt.Println(tid)
	fmt.Println(content)
	fmt.Println(creatorID)

	app.messages.InsertMessageInThread(tid, content, creatorID)

	http.Redirect(w, r, fmt.Sprintf("/forum/thread?thread=%s", tt), http.StatusSeeOther)
}

func (app *application) createSubject(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		app.clientError(w, http.StatusMethodNotAllowed)
		return
	}

	var form createForm

	err := r.ParseForm()
	if err != nil {
		app.serverError(w, err)
		return
	}

	app.formDecoder.Decode(form, r.Form)

	fmt.Println(r.FormValue("subject"))

	content := r.FormValue("subject")

	app.messages.InsertSubject(content)

	http.Redirect(w, r, fmt.Sprintf("/forum/subjects"), http.StatusSeeOther)

}

func (app *application) createThread(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		app.clientError(w, http.StatusMethodNotAllowed)
		return
	}

	err := r.ParseForm()
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	threadTitle := r.FormValue("threadtitle")
	subjectId, err := strconv.Atoi(r.FormValue("subjectid"))
	if err != nil {
		app.serverError(w, err)
		return
	}

	app.messages.InsertThreadInSubject(subjectId, threadTitle)
	http.Redirect(w, r, fmt.Sprintf("/forum/subject?subject=%s", r.FormValue("subjecttitle")), http.StatusSeeOther)
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
