package session

import (
	"ContactBook/internal/contactBook"
	"sync"
)

type SessionManager struct {
	mu      sync.Mutex
	session map[string]*contactBook.ContactManager
}

func NewSessionManager() *SessionManager {
	return &SessionManager{
		session: make(map[string]*contactBook.ContactManager),
	}
}

func (s *SessionManager) GetSession(token string) *contactBook.ContactManager {
	sm, ok := s.session[token]
	if ok {
		return sm
	}
	s.mu.Lock()
	defer s.mu.Unlock()
	newContactManager := contactBook.NewContactManager()
	s.session[token] = newContactManager
	return newContactManager
}
