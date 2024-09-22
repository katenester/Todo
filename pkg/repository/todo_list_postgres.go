package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	todo "github.com/katenester/Todo"
)

type TodoListPostgres struct {
	db *sqlx.DB
}

func NewTodoListPostgres(db *sqlx.DB) *TodoListPostgres {
	return &TodoListPostgres{db: db}
}

func (t *TodoListPostgres) Create(userId int, list todo.TodoList) (int, error) {
	tx, err := t.db.Begin()
	if err != nil {
		return 0, err
	}
	var id int
	createListQuery := fmt.Sprintf("INSERT INTO %s (title,description) VALUES ($1,$2) RETURNING id", todoListsTable)
	if err := tx.QueryRow(createListQuery, list.Title, list.Description).Scan(&id); err != nil {
		tx.Rollback()
		return 0, err
	}
	createUserListQuery := fmt.Sprintf("INSERT INTO %s (user_id,list_id) VALUES ($1,$2)", usersListsTable)
	if _, err := tx.Exec(createUserListQuery, userId, id); err != nil {
		tx.Rollback()
		return 0, err
	}
	return id, tx.Commit()
}
func (t *TodoListPostgres) GetAll(userId int) ([]todo.TodoList, error) {
	var lists []todo.TodoList
	query := fmt.Sprintf("SELECT t1.id, t1.title,t1.description FROM %s t1 INNER JOIN %s t2 on t1.id=t2.list_id WHERE t2.user_id=$1", todoListsTable, usersListsTable)
	err := t.db.Select(&lists, query, userId)
	return lists, err
}

func (t *TodoListPostgres) GetById(userId int, listId int) (todo.TodoList, error) {
	var list todo.TodoList
	query := fmt.Sprintf(`SELECT t1.id, t1.title,t1.description FROM %s t1 inner join %s t2
		on t1.id=t2.list_id where t1.id=$1 and t2.user_id=$2`, todoListsTable, usersListsTable)
	err := t.db.Get(&list, query, listId, userId)
	return list, err
}
