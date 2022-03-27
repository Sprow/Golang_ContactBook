package db

import (
	"database/sql"
	_ "embed"
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
	err = migrate(db)
	if err != nil {
		log.Println(err)
	}
	log.Println("migrate complete")
	return db, nil
}

//go:embed migrate.sql
var migrationSql string

func migrate(db Database) error {
	_, err := db.Conn.Exec(migrationSql)
	return err
}
