package models

import (
	"database/sql"
	"fmt"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type Message struct {
	ThreadID    int
	Content     string
	CreatorID   int
	DateCreated time.Time
	MessageID   int
}

type Thread struct {
	ThreadID    int
	SubjectID   int
	DateCreated time.Time
	Title       string
}

type Subject struct {
	SubjectID int
	Title     string
}

type MessageModel struct {
	DB *sql.DB
}

type ThreadModel struct {
	DB *sql.DB
}

type SubjectModel struct {
	DB *sql.DB
}

type User struct {
	Email    string
	Password string
	Name     string
}

func (m *MessageModel) InsertMessageInThread(threadId int, content string, creatorId int) (int, error) {
	query := `INSERT INTO messages (thread_id, content, creator_id, date_created)
	VALUES(?, ?, ?, UTC_TIMESTAMP())`

	result, err := m.DB.Exec(query, threadId, content, creatorId)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

func (m *MessageModel) InsertThreadInSubject(subjectId int, title string) (int, error) {
	query := `INSERT INTO threads (title, subject_id, date_created) VALUES(?,?,UTC_TIMESTAMP())`

	result, err := m.DB.Exec(query, title, subjectId)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

func (m *MessageModel) InsertSubject(title string) (int, error) {
	query := `INSERT INTO subjects (title) VALUES(?)`

	result, err := m.DB.Exec(query, title)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

func (m *MessageModel) RegisterNewUser(name string, password string, email string) (int64, error) {
	query := `INSERT INTO users (name, email, hashed_password, created) VALUES(?,?,?,UTC_TIMESTAMP())`

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	if err != nil {
		return 0, err
	}

	result, err := m.DB.Exec(query, name, email, hashedPassword)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (m *MessageModel) GetMessagesInThread(threadId int) ([]*Message, error) {
	query := `SELECT content, creator_id, date_created, thread_id, message_id FROM messages WHERE
	thread_id = ?`

	rows, err := m.DB.Query(query, threadId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	messages := []*Message{}

	for rows.Next() {
		messag := &Message{}

		err := rows.Scan(&messag.Content, &messag.CreatorID, &messag.DateCreated, &messag.ThreadID, &messag.MessageID)
		if err != nil {
			return nil, err
		}
		messages = append(messages, messag)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return messages, nil
}

func (m *MessageModel) GetAllSubjects() ([]*Subject, error) {
	query := `SELECT id, title FROM subjects`

	results, err := m.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer results.Close()

	subjects := []*Subject{}

	for results.Next() {
		subject := &Subject{}
		err := results.Scan(&subject.SubjectID, &subject.Title)
		if err != nil {
			return nil, err
		}
		subjects = append(subjects, subject)
	}
	return subjects, nil
}

func (m *MessageModel) GetThreadsInSubject(subjectId int) ([]*Thread, error) {
	query := `SELECT title, date_created, thread_id, subject_id FROM threads WHERE subject_id=?`

	results, err := m.DB.Query(query, subjectId)
	if err != nil {
		return nil, err
	}
	defer results.Close()

	threads := []*Thread{}

	for results.Next() {
		thread := &Thread{}
		err := results.Scan(&thread.Title, &thread.DateCreated, &thread.ThreadID, &thread.SubjectID)
		if err != nil {
			return nil, err
		}
		threads = append(threads, thread)
	}
	return threads, nil
}

func (m *MessageModel) GetThreadId(title string) (int, error) {
	query := `SELECT thread_id FROM threads WHERE title = ?`

	var titleId *int

	id := m.DB.QueryRow(query, title)
	err := id.Scan(&titleId)
	if err != nil {
		return 0, err
	}
	return *titleId, nil
}

func (m *MessageModel) GetSubjectId(title string) (int, error) {
	query := `SELECT id FROM subjects WHERE title = ?`

	var titleId *int
	fmt.Println(title)
	id := m.DB.QueryRow(query, title)
	err := id.Scan(&titleId)
	if err != nil {
		return 0, err
	}

	return *titleId, nil
}

func (m *MessageModel) Authenticate(name, password string) (int, error) {
	var id int
	var hashedPassword []byte

	query := "SELECT user_id, hashed_password FROM users WHERE name = ?"

	err := m.DB.QueryRow(query, name).Scan(&id, &hashedPassword)
	if err != nil {
		fmt.Println("ERROR IN DB QUERY ROW")
		return 0, err
	}

	err = bcrypt.CompareHashAndPassword(hashedPassword, []byte(password))
	if err != nil {
		return 0, err
	}

	return id, nil
}
