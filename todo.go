package todo

import (
	"errors"
	"time"
)

type TodoCategory struct {
	Id    int    `json:"id" db:"id"`
	Title string `json:"title" db:"title" binding:"required"`
}

type TodoItem struct {
	Id          int       `json:"id" db:"id"`
	Title       string    `json:"title" db:"title" binding:"required"`
	Description string    `json:"description" db:"description" binding:"required"`
	IsCompleted bool      `json:"is_completed" db:"is_completed"`
	IsImportant bool      `json:"is_important" db:"is_important"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
}

type UpdateTodoCategoryInput struct {
	Title *string `json:"title"`
}

func (i UpdateTodoCategoryInput) Validate() error {
	if i.Title == nil {
		return errors.New("update structure has no values")
	}
	return nil
}

type UpdateTodoItemInput struct {
	Title       *string `json:"title"`
	Description *string `json:"description"`
	IsCompleted *bool   `json:"is_completed"`
	IsImportant *bool   `json:"is_important"`
}

func (i UpdateTodoItemInput) Validate() error {
	if i.Title == nil && i.Description == nil && i.IsCompleted == nil && i.IsImportant == nil {
		return errors.New("update structure has no values")
	}
	return nil
}
