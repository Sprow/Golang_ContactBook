package session

import (
	"ContactBook/internal/contactBook"
	"ContactBook/internal/db"
	"github.com/google/uuid"
	"log"
	"time"
)

type SessionManager struct {
	db      db.Database
	session map[string]*contactBook.ContactManager
}

func NewSessionManager(db db.Database) *SessionManager {
	return &SessionManager{
		db:      db,
		session: make(map[string]*contactBook.ContactManager),
	}
}

func (s *SessionManager) CreateNewUserSession() (id string, err error) {
	newUserId := uuid.New()
	datatime := time.Now()

	res, err := s.db.Conn.Exec(
		"INSERT INTO users (id, created_at) VALUES ($1, $2)", newUserId, datatime)

	log.Println("result of creating new user>>>", res)
	if err != nil {
		return newUserId.String(), err
	}
	newToken := uuid.New()
	res, err = s.db.Conn.Exec(
		"INSERT INTO sessions (token, user_id) VALUES ($1, $2)", newToken, newUserId)

	if err != nil {
		return newUserId.String(), err
	}

	return newUserId.String(), nil
}

func (s *SessionManager) GetSession(token string) *contactBook.ContactManager {
	sm, ok := s.session[token]
	if ok {
		return sm
	}

	id, _ := s.CreateNewUserSession()
	newContactManager := contactBook.NewContactManager(s.db, id)
	log.Println("newContactManager id >>>", id)
	s.session[token] = newContactManager
	return newContactManager
}
