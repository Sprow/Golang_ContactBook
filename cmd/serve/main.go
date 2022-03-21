package main

import (
	"ContactBook/cmd/serve/handler"
	"ContactBook/internal/session"
	"fmt"
	"github.com/go-chi/chi/v5"
	"net/http"
)

func main() {
	//cm := contactBook.NewContactManager() // change contactManger to sessionManager in handler
	sm := session.NewSessionManager()
	h := handler.NewHandler(sm)
	router := chi.NewRouter()
	h.Register(router)
	err := http.ListenAndServe(Port, router)
	if err != nil {
		fmt.Println(err)
	}
}
