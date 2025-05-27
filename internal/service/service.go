package service

import (
	todo "github.com/al1enn/go_todo_app"
	"github.com/al1enn/go_todo_app/internal/repository"
)

type Authorization interface {
	CreateUser(user todo.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
}

type TodoCategory interface {
	Create(userId int, category todo.TodoCategory) (int, error)
	GetAll(userId int) ([]todo.TodoCategory, error)
	GetById(userId, categoryId int) (todo.TodoCategory, error)
	Delete(userId, categoryId int) error
	Update(userId, categoryId int, input todo.UpdateTodoCategoryInput) error
}

type TodoItem interface {
	Create(userId, categoryId int, item todo.TodoItem) (int, error)
	GetAll(userId int) ([]todo.TodoItem, error)
	GetById(userId, itemId int) (todo.TodoItem, error)
	Delete(userId, itemId int) error
	Update(userId, itemId int, input todo.UpdateTodoItemInput) error
}

type Service struct {
	Authorization
	TodoCategory
	TodoItem
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		TodoCategory:  NewTodoCategoryService(repos.TodoCategory),
		TodoItem:      NewTodoItemService(repos.TodoItem, repos.TodoCategory),
	}
}
