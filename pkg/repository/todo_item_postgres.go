package repository

import (
	"fmt"
	"strings"

	todo "github.com/al1enn/go_todo_app"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
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
	createItemQuery := fmt.Sprintf("INSERT INTO %s (title, description,is_completed,is_important) VALUES ($1, $2, $3, $4) RETURNING id",
		todoItemsTable)
	row := tx.QueryRow(createItemQuery, item.Title, item.Description, item.IsCompleted, item.IsImportant)
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

func (r *TodoItemPostgres) GetById(userId, id int) (todo.TodoItem, error) {
	var item todo.TodoItem
	query := fmt.Sprintf(`SELECT ti.id, ti.title, ti.description, ti.is_completed 
		FROM %s ti 
		INNER JOIN %s tic ON tic.todo_item_id = ti.id
		INNER JOIN %s tc ON tc.id = tic.todo_category_id  
		INNER JOIN %s utc ON utc.todo_category_id = tc.id
		WHERE utc.user_id = $1 AND ti.id = $2`,
		todoItemsTable, todoItemsCategoriesTable, todoCategoriesTable, usersTodoCategoriesTable)
	if err := r.db.Get(&item, query, userId, id); err != nil {
		return item, err
	}
	return item, nil
}

func (r *TodoItemPostgres) Delete(userId, id int) error {
	query := fmt.Sprintf(`DELETE FROM %s ti USING %s tic, %s tc, %s utc 
		WHERE ti.id = tic.todo_item_id AND tic.todo_category_id = tc.id AND tc.id = utc.todo_category_id AND utc.user_id = $1 AND ti.id = $2`,
		todoItemsTable, todoItemsCategoriesTable, todoCategoriesTable, usersTodoCategoriesTable)
	_, err := r.db.Exec(query, userId, id)
	return err
}

func (r *TodoItemPostgres) Update(userId, itemId int, input todo.UpdateTodoItemInput) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if input.Title != nil {
		setValues = append(setValues, fmt.Sprintf("title=$%d", argId))
		args = append(args, *input.Title)
		argId++
	}
	if input.Description != nil {
		setValues = append(setValues, fmt.Sprintf("description=$%d", argId))
		args = append(args, *input.Description)
		argId++
	}
	if input.IsCompleted != nil {
		setValues = append(setValues, fmt.Sprintf("is_completed=$%d", argId))
		args = append(args, *input.IsCompleted)
		argId++
	}
	if input.IsImportant != nil {
		setValues = append(setValues, fmt.Sprintf("is_important=$%d", argId))
		args = append(args, *input.IsImportant)
		argId++
	}
	setQuery := strings.Join(setValues, ", ")
	args = append(args, userId, itemId)

	query := fmt.Sprintf(`UPDATE %s ti SET %s FROM %s tic, %s tc, %s utc 
	                    WHERE ti.id = tic.todo_item_id AND tic.todo_category_id = tc.id 
						AND tc.id = utc.todo_category_id AND utc.user_id = $%d AND ti.id = $%d`,
		todoItemsTable, setQuery, todoItemsCategoriesTable, todoCategoriesTable, usersTodoCategoriesTable, argId, argId+1)

	_, err := r.db.Exec(query, args...)

	logrus.Debugf("updateQuery: %s", query)
	logrus.Debugf("args: %s", args)

	return err
}
