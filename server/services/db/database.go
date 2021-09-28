package db

import (
	"database/sql"
  "fmt"
	"os"

	_ "github.com/lib/pq" // Needed for the postgres driver
)


// DB is the struct that we're going to use to implement all of our
// Datasbase interfaces; All of the methods defined on each of our
// interfaces will be implemented on this DB struct
type DB struct {
  *sql.DB
}


// DatabaseManager combines all of the database interactions into one
type DatabaseManager interface {
  CardManager
  SetManager
}


// New initializes a new postgres database connection and attaches
// said connection to our DB struct, which we can then call all of
// the methods described by the our varies Database interfaces
func New() (*DB, error) {
  connStr := fmt.Sprintf(
    "user=%s password=%s dbname=%s host=%s port=%s",
    os.Getenv("USER_NAME"),
    os.Getenv("USER_PASSWORD"),
    os.Getenv("DB_NAME"),
    os.Getenv("DB_HOST"),
    os.Getenv("DB_PORT"),
  )

  db, err := sql.Open("postgres", connStr)
  if err != nil {
    return nil, err
  }
  if err = db.Ping(); err != nil {
    return nil, err
  }
  return &DB{db}, nil
}
