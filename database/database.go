package database

import (
	"database/sql"
	"fmt"
	"github.com/rohit0700/go-api-framework/config"

	"github.com/Masterminds/squirrel"
	_ "github.com/lib/pq"
)

var db *sql.DB

func InitDatabase(cfg *config.Config) (*sql.DB, error) {
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", cfg.Database.User, cfg.Database.Password, cfg.Database.Host, cfg.Database.Port, cfg.Database.Name)

	conn, err := sql.Open(cfg.Database.Driver, connectionString)
	if err != nil {
		return nil, err
	}

	db = conn
	return db, nil
}

func FetchUser(userID int) (string, error) {
	// Use Squirrel for fetching user information
	// Example query: SELECT name FROM users WHERE id = ?
	query := squirrel.Select("name").From("users").Where(squirrel.Eq{"id": userID})

	// Execute query
	row := query.RunWith(db).QueryRow()

	var name string
	err := row.Scan(&name)
	if err != nil {
		return "", err
	}

	return name, nil
}

func InsertUser(name string) error {
	// Use Squirrel for inserting user information
	// Example query: INSERT INTO users (name) VALUES (?)
	query := squirrel.Insert("users").Columns("name").Values(name)

	_, err := query.RunWith(db).Exec()
	if err != nil {
		return err
	}

	return nil
}
