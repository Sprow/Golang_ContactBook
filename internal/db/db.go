package db

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

const (
	HOST = "database"
	PORT = 5432
)

type Config struct {
	Username string
	Password string
	DBName   string
}

var ErrNoMatch = fmt.Errorf("no matching record")

type Database struct {
	Conn *sql.DB
}

func Initialize(cfg Config) (Database, error) {
	log.Println("Initializing")
	db := Database{}
	connStr := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable",
		cfg.Username, cfg.Password, HOST, PORT, cfg.DBName)
	conn, err := sql.Open("postgres", connStr)
	if err != nil {
		return db, err
	}

	db.Conn = conn
	err = db.Conn.Ping()
	if err != nil {
		return db, err
	}
	log.Println("Database connection established")
	return db, nil
}
