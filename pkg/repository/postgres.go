package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

const (
	usersTable               = "users"
	todoItemsTable           = "todo_items"
	todoCategoriesTable      = "todo_categories"
	usersTodoCategoriesTable = "user_todo_categories"
	todoItemsCategoriesTable = "todo_item_categories"
)

type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SSLMode  string
}

func NewPostgresDB(config Config) (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s", config.Host, config.Port,
		config.Username, config.DBName, config.Password, config.SSLMode))
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, nil
}
