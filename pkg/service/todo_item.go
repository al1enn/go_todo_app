package service

import (
	todo "github.com/al1enn/go_todo_app"
	"github.com/al1enn/go_todo_app/pkg/repository"
)

type TodoItemService struct {
	repo         repository.TodoItem
	categoryRepo repository.TodoCategory
}

func NewTodoItemService(repo repository.TodoItem) *TodoItemService {
	return &TodoItemService{
		repo: repo,
	}
}

func (s *TodoItemService) Create(userId int, item todo.TodoItem) (int, error) {
	_, err := s.categoryRepo.GetById(userId, item.CategoryId)
	if err != nil {
		return 0, err
	}
	return s.repo.Create(item)
}
