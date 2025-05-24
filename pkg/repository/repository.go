package repository

import (
	todo "github.com/al1enn/go_todo_app"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user todo.User) (int, error)
	GetUser(username, password string) (todo.User, error)
}

type TodoCategory interface {
	Create(userId int, category todo.TodoCategory) (int, error)
	GetAll(userId int) ([]todo.TodoCategory, error)
	GetById(userId, id int) (todo.TodoCategory, error)
	Delete(userId, id int) error
	Update(userId, id int, input todo.UpdateTodoCategoryInput) error
}

type TodoItem interface {
	Create(categoryId int, item todo.TodoItem) (int, error)
	GetAll(userId int) ([]todo.TodoItem, error)
	GetById(userId, id int) (todo.TodoItem, error)
	Delete(userId, id int) error
}

type Repository struct {
	Authorization
	TodoCategory
	TodoItem
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		TodoCategory:  NewTodoCategoryPostgres(db),
		TodoItem:      NewTodoItemPostgres(db),
	}
}
