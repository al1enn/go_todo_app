package repository

import (
	"fmt"
	"strings"

	todo "github.com/al1enn/go_todo_app"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type TodoCategoryPostgres struct {
	db *sqlx.DB
}

func NewTodoCategoryPostgres(db *sqlx.DB) *TodoCategoryPostgres {
	return &TodoCategoryPostgres{
		db: db,
	}
}

func (repo *TodoCategoryPostgres) Create(userId int, category todo.TodoCategory) (int, error) {
	tx, err := repo.db.Begin()
	if err != nil {
		return 0, err
	}

	var todoCategoryId int
	createTodoCategoryQuery := fmt.Sprintf("INSERT INTO %s (title) VALUES ($1) RETURNING id", todoCategoriesTable)
	row := tx.QueryRow(createTodoCategoryQuery, category.Title)
	if err := row.Scan(&todoCategoryId); err != nil {
		tx.Rollback()
		return 0, err
	}
	createUsersTodoCategoryQuery := fmt.Sprintf("INSERT INTO %s (user_id, todo_category_id) VALUES ($1, $2)", usersTodoCategoriesTable)
	_, err = tx.Exec(createUsersTodoCategoryQuery, userId, todoCategoryId)
	if err != nil {
		tx.Rollback()
		return 0, err
	}
	return todoCategoryId, tx.Commit()
}

func (repo *TodoCategoryPostgres) GetAll(userId int) ([]todo.TodoCategory, error) {
	var todoCategories []todo.TodoCategory
	query := fmt.Sprintf("SELECT tl.id, tl.title FROM %s tl INNER JOIN %s utl ON tl.id = utl.todo_category_id WHERE utl.user_id = $1",
		todoCategoriesTable, usersTodoCategoriesTable)
	err := repo.db.Select(&todoCategories, query, userId)
	return todoCategories, err
}

func (repo *TodoCategoryPostgres) GetById(userId, id int) (todo.TodoCategory, error) {
	var todoCategory todo.TodoCategory
	query := fmt.Sprintf(`SELECT tl.id, tl.title FROM %s tl 
						INNER JOIN %s utl ON tl.id = utl.todo_category_id WHERE utl.user_id = $1 AND tl.id=$2`,
		todoCategoriesTable, usersTodoCategoriesTable)
	err := repo.db.Get(&todoCategory, query, userId, id)
	return todoCategory, err
}

func (repo *TodoCategoryPostgres) Delete(userId, id int) error {
	query := fmt.Sprintf("DELETE FROM %s tl USING %s utl WHERE tl.id = $1 AND utl.user_id = $2 AND tl.id = utl.todo_category_id",
		todoCategoriesTable, usersTodoCategoriesTable)
	_, err := repo.db.Exec(query, id, userId)
	return err
}

func (r *TodoCategoryPostgres) Update(userId, todoCategoryId int, input todo.UpdateTodoCategoryInput) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if input.Title != nil {
		setValues = append(setValues, fmt.Sprintf("title=$%d", argId))
		args = append(args, *input.Title)
		argId++
	}

	setQuery := strings.Join(setValues, ", ")

	query := fmt.Sprintf("UPDATE %s tc SET %s FROM %s utc WHERE tc.id = utc.todo_category_id AND utc.todo_category_id=$%d AND utc.user_id=$%d",
		todoCategoriesTable, setQuery, usersTodoCategoriesTable, argId, argId+1)
	args = append(args, todoCategoryId, userId)

	logrus.Debugf("updateQuery: %s", query)
	logrus.Debugf("args: %s", args)

	_, err := r.db.Exec(query, args...)
	return err
}
