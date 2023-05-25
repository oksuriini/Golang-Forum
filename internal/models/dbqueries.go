package models

import (
	"database/sql"
	"time"
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

func (m *MessageModel) Insert(threadId int, content string, creatorId int) (int, error) {
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

func (t *ThreadModel) Insert(threadId int, subjectId int) (int, error) {
	return 0, nil
}

func (s *SubjectModel) Insert(subjectId int, title string) (int, error) {
	return 0, nil
}

func (m *MessageModel) Get(threadId int) ([]*Message, error) {
	query := `SELECT content, creator_id FROM messages WHERE
	thread_id = ?`

	result, err := m.DB.Query(query, threadId)
	if err != nil {
		return nil, err
	}

	if result != nil {
		return nil, nil
	}
	return nil, nil
}

func (t *ThreadModel) Get(subjectId int) (*Thread, error) {
	return nil, nil
}

func (s *SubjectModel) Get() (*Subject, error) {
	return nil, nil
}
