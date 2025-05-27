package service

import (
	todo "github.com/al1enn/go_todo_app"
	"github.com/al1enn/go_todo_app/internal/repository"
)

type TodoItemService struct {
	repo         repository.TodoItem
	categoryRepo repository.TodoCategory
}

func NewTodoItemService(repo repository.TodoItem, categoryRepo repository.TodoCategory) *TodoItemService {
	return &TodoItemService{
		repo:         repo,
		categoryRepo: categoryRepo,
	}
}

func (s *TodoItemService) Create(userId, categoryId int, item todo.TodoItem) (int, error) {
	_, err := s.categoryRepo.GetById(userId, categoryId)
	if err != nil {
		return 0, err
	}
	return s.repo.Create(categoryId, item)
}

func (s *TodoItemService) GetAll(userId int) ([]todo.TodoItem, error) {
	return s.repo.GetAll(userId)
}

func (s *TodoItemService) GetById(userId, id int) (todo.TodoItem, error) {
	return s.repo.GetById(userId, id)
}

func (s *TodoItemService) Delete(userId, id int) error {
	return s.repo.Delete(userId, id)
}

func (s *TodoItemService) Update(userId, id int, input todo.UpdateTodoItemInput) error {
	if err := input.Validate(); err != nil {
		return err
	}
	return s.repo.Update(userId, id, input)
}
