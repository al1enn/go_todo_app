package repository

import (
	"fmt"

	todo "github.com/al1enn/go_todo_app"
	"github.com/jmoiron/sqlx"
)

type TodoItemPostgres struct {
	db *sqlx.DB
}

func NewTodoItemPostgres(db *sqlx.DB) *TodoItemPostgres {
	return &TodoItemPostgres{
		db: db,
	}
}

func (r *TodoItemPostgres) Create(item todo.TodoItem) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}
	var itemId int
	createItemQuery := fmt.Sprintf("INSERT INTO %s (title, description, todo_category_id) VALUES ($1, $2, $3) RETURNING id",
		todoItemsTable)
	row := tx.QueryRow(createItemQuery, item.Title, item.Description, item.CategoryId)
	if err := row.Scan(&itemId); err != nil {
		tx.Rollback()
		return 0, err
	}
	createTodoItemsCategoryQuery := fmt.Sprintf("INSERT INTO %s (todo_item_id, todo_category_id) VALUES ($1, $2)",
		todoItemsCategoriesTable)
	_, err = tx.Exec(createTodoItemsCategoryQuery, itemId, item.CategoryId)
	if err != nil {
		tx.Rollback()
		return 0, err
	}
	return itemId, tx.Commit()
}
