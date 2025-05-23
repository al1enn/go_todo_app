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

func (r *TodoItemPostgres) Create(categoryId int, item todo.TodoItem) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}
	var itemId int
	createItemQuery := fmt.Sprintf("INSERT INTO %s (title, description) VALUES ($1, $2) RETURNING id",
		todoItemsTable)
	row := tx.QueryRow(createItemQuery, item.Title, item.Description)
	if err := row.Scan(&itemId); err != nil {
		tx.Rollback()
		return 0, err
	}
	createTodoItemsCategoryQuery := fmt.Sprintf("INSERT INTO %s (todo_item_id, todo_category_id) VALUES ($1, $2)",
		todoItemsCategoriesTable)
	_, err = tx.Exec(createTodoItemsCategoryQuery, itemId, categoryId)
	if err != nil {
		tx.Rollback()
		return 0, err
	}
	return itemId, tx.Commit()
}

func (r *TodoItemPostgres) GetAll(userId int) ([]todo.TodoItem, error) {
	var items []todo.TodoItem
	query := fmt.Sprintf(`SELECT ti.id, ti.title, ti.description, ti.is_completed 
		FROM %s ti 
		INNER JOIN %s tic ON tic.todo_item_id = ti.id
		INNER JOIN %s tc ON tc.id = tic.todo_category_id  
		INNER JOIN %s utc ON utc.todo_category_id = tc.id
		WHERE utc.user_id = $1`,
		todoItemsTable, todoItemsCategoriesTable, todoCategoriesTable, usersTodoCategoriesTable)
	if err := r.db.Select(&items, query, userId); err != nil {
		return nil, err
	}
	return items, nil
}
