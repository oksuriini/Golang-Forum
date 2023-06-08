package main

import (
	"errors"
	"fmt"
	"html/template"
	"net/http"
	"strconv"

	"github.com/go-playground/form/v4"
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
	Subject     string `form:"subject"`
	SubjectID   string `form:"subjectid"`
}

type userForm struct {
	Name     string
	Email    string
	Password string
}

// Homepage handler
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

// Forum mainpage
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

// ----------------------------------------------------------------------
// GET Handlers

// Fetches messages of specific thread
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

// Fetches all threads under a subject
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

// Fetches all subjects
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

// Fetches user registration form
func (app *application) registerUser(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/forum/registrar" {
		app.notFound(w)
		return
	}

	files := []string{
		"./ui/html/base.tmpl.html",
		"./ui/html/partials/footer.tmpl.html",
		"./ui/html/partials/nav.tmpl.html",
		"./ui/html/pages/register.tmpl.html",
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

// Fetches user login form
func (app *application) loginUser(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/forum/login" {
		app.notFound(w)
		return
	}

	files := []string{
		"./ui/html/base.tmpl.html",
		"./ui/html/partials/footer.tmpl.html",
		"./ui/html/partials/nav.tmpl.html",
		"./ui/html/pages/login.tmpl.html",
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

// ----------------------------------------------------------------------
// POST Handlers

// Creates a message under a specific thread with POST method lives under getThreadMessages handler
func (app *application) createMessage(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		app.clientError(w, http.StatusMethodNotAllowed)
		return
	}

	var form createForm

	app.decodePostForm(r, &form)

	tid, err := app.messages.GetThreadId(form.ThreadTitle)
	if err != nil {
		app.serverError(w, err)
		return
	}

	creatorInt, err := strconv.Atoi(form.CreatorID)
	if err != nil {
		app.serverError(w, err)
		return
	}

	app.messages.InsertMessageInThread(tid, form.Content, creatorInt)

	http.Redirect(w, r, fmt.Sprintf("/forum/thread?thread=%s", form.ThreadTitle), http.StatusSeeOther)
}

// Creates a subject on main forum with POST method lives under getSubjects handler
func (app *application) createSubject(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		app.clientError(w, http.StatusMethodNotAllowed)
		return
	}

	var form createForm

	app.decodePostForm(r, &form)

	app.messages.InsertSubject(form.Subject)

	http.Redirect(w, r, fmt.Sprintf("/forum/subjects"), http.StatusSeeOther)
}

// Create a thread under a specified subject with POST method lives under getThreads
func (app *application) createThread(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		app.clientError(w, http.StatusMethodNotAllowed)
		return
	}

	var form createForm

	app.decodePostForm(r, &form)

	subjectId, err := strconv.Atoi(form.SubjectID)
	if err != nil {
		app.serverError(w, err)
		return
	}

	app.messages.InsertThreadInSubject(subjectId, form.ThreadTitle)
	http.Redirect(w, r, fmt.Sprintf("/forum/subject?subject=%s", form.Subject), http.StatusSeeOther)
}

// Create a user POST method lives under registerUser handler
func (app *application) registerUserPost(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		app.clientError(w, http.StatusMethodNotAllowed)
		return
	}

	err := r.ParseForm()
	if err != nil {
		app.serverError(w, err)
		return
	}

	_, err = app.messages.RegisterNewUser(r.FormValue("name"), r.FormValue("password"), r.FormValue("email"))
	if err != nil {
		app.serverError(w, err)
		return
	}
	http.Redirect(w, r, "/forum", http.StatusSeeOther)
}

// Login a user with POST method lives under loginUser handler
func (app *application) loginUserPost(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		app.clientError(w, http.StatusMethodNotAllowed)
		return
	}

	err := r.ParseForm()
	if err != nil {
		app.serverError(w, err)
		return
	}

	id, err := app.messages.Authenticate(r.FormValue("name"), r.FormValue("password"))

	if id == 0 {
		app.infoLogger.Println("Authentication failed for user")
		http.Redirect(w, r, "/forum/login", http.StatusBadRequest)
		return
	}

	fmt.Printf("Authenticated user id:%d", id)

	http.Redirect(w, r, "/forum", http.StatusSeeOther)
}

func (app *application) logoutUserPost(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		app.clientError(w, http.StatusBadRequest)
		return
	}

}

func (app *application) decodePostForm(req *http.Request, dst any) error {

	err := req.ParseForm()
	if err != nil {
		return err
	}

	err = app.formDecoder.Decode(dst, req.PostForm)
	if err != nil {
		var invalidDecoderError *form.InvalidDecoderError

		if errors.As(err, &invalidDecoderError) {
			panic(err)
		}
		return err
	}
	return nil
}
