package main

import (
	"ContactBook/cmd/serve/handler"
	"ContactBook/internal/db"
	"ContactBook/internal/session"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

// TODO
// postgres
// try to use new generic to clone map

func main() {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	environmentPath := filepath.Join(dir, ".env")
	err = godotenv.Load(environmentPath) // load .env
	cfg := db.Config{
		Username: os.Getenv("POSTGRES_USER"),
		Password: os.Getenv("POSTGRES_PASSWORD"),
		DBName:   os.Getenv("POSTGRES_DB"),
	}
	database, err := db.Initialize(cfg)
	if err != nil {
		log.Fatalf("failed to initialize database: %s", err.Error())
	}
	fmt.Println(database.Conn.Ping())

	//cm := contactBook.NewContactManager() // change contactManger to sessionManager in handler
	sm := session.NewSessionManager()
	h := handler.NewHandler(sm)
	router := chi.NewRouter()
	router.Use(middleware.Logger)
	h.Register(router)
	err = http.ListenAndServe(Port, router)
	if err != nil {
		fmt.Println(err)
	}
}
