package handler

import (
	"ContactBook/internal/contactBook"
	"ContactBook/internal/session"
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/rs/xid"
	"net/http"
	"time"
)

type Handler struct {
	//contactManager *contactBook.ContactManager
	sessionManager *session.SessionManager
}

//func NewHandler(contactManager *contactBook.ContactManager) *Handler {
//	return &Handler{
//		contactManager: contactManager,
//	}
//}

func NewHandler(sessionManager *session.SessionManager) *Handler {
	return &Handler{
		sessionManager: sessionManager,
	}
}

func (h *Handler) Register(r *chi.Mux) {
	r.Post("/add_contact", h.sessionMiddleware(h.addContact))
	r.Post("/remove_contact", h.sessionMiddleware(h.removeContact))
	r.Post("/update_contact", h.sessionMiddleware(h.updateContact))
	r.Get("/", h.sessionMiddleware(h.getAllContacts))
}

func (h *Handler) addContact(w http.ResponseWriter, r *http.Request) {
	contactManager := h.sessionManager.GetSession(h.getToken(r))
	d := json.NewDecoder(r.Body)
	var contact contactBook.Contact
	err := d.Decode(&contact)
	if err != nil {
		fmt.Println(err)
	}
	//h.contactManager.AddContact(contact)
	contactManager.AddContact(contact)
}

type Index struct {
	Id int `json:"id"`
}

func (h *Handler) removeContact(w http.ResponseWriter, r *http.Request) {
	contactManager := h.sessionManager.GetSession(h.getToken(r))
	d := json.NewDecoder(r.Body)
	var id Index
	err := d.Decode(&id)
	if err != nil {
		fmt.Println(err)
	}
	contactManager.RemoveContact(id.Id)
}

func (h *Handler) updateContact(w http.ResponseWriter, r *http.Request) {
	contactManager := h.sessionManager.GetSession(h.getToken(r))
	d := json.NewDecoder(r.Body)
	var updateObj contactBook.UpdateContactDTO
	err := d.Decode(&updateObj)
	if err != nil {
		fmt.Println(err)
	}
	contactManager.UpdateContact(updateObj)
}

func (h *Handler) getAllContacts(w http.ResponseWriter, r *http.Request) {
	contactManager := h.sessionManager.GetSession(h.getToken(r))
	encoder := json.NewEncoder(w)
	data := contactManager.GetAllContacts()
	err := encoder.Encode(data)
	if err != nil {
		fmt.Println(err)
	}
}

func (h *Handler) getToken(r *http.Request) string {
	c, err := r.Cookie("token")
	if err != nil {
		return ""
	}
	return c.Value
}

func (h *Handler) sessionMiddleware(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		t := h.getToken(r)
		if t == "" {
			cookie := http.Cookie{
				Name:     "token",
				Value:    xid.New().String(),             //new token
				Expires:  time.Now().Add(24 * time.Hour), //user lose his contact book after 24h
				HttpOnly: true,
			}
			http.SetCookie(w, &cookie)

			r.AddCookie(&cookie)
		}
		handler(w, r)
	}
}
