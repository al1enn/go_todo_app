package service

import (
	todo "github.com/al1enn/go_todo_app"
	"github.com/al1enn/go_todo_app/internal/repository"
)

type TodoCategoryService struct {
	repo repository.TodoCategory
}

func NewTodoCategoryService(repo repository.TodoCategory) *TodoCategoryService {
	return &TodoCategoryService{
		repo: repo,
	}
}

func (s *TodoCategoryService) Create(userId int, category todo.TodoCategory) (int, error) {
	return s.repo.Create(userId, category)
}

func (s *TodoCategoryService) GetAll(userId int) ([]todo.TodoCategory, error) {
	return s.repo.GetAll(userId)
}

func (s *TodoCategoryService) GetById(userId, id int) (todo.TodoCategory, error) {
	return s.repo.GetById(userId, id)
}

func (s *TodoCategoryService) Delete(userId, id int) error {
	return s.repo.Delete(userId, id)
}

func (s *TodoCategoryService) Update(userId, id int, input todo.UpdateTodoCategoryInput) error {
	if err := input.Validate(); err != nil {
		return err
	}
	return s.repo.Update(userId, id, input)
}
