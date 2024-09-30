package todo_list

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	todo "github.com/katenester/Todo/internal/models"
	"github.com/katenester/Todo/internal/repository/postgres/config"
	"github.com/sirupsen/logrus"
	"strings"
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
	createListQuery := fmt.Sprintf("INSERT INTO %s (title,description) VALUES ($1,$2) RETURNING id", config.TodoListsTable)
	if err := tx.QueryRow(createListQuery, list.Title, list.Description).Scan(&id); err != nil {
		tx.Rollback()
		return 0, err
	}
	createUserListQuery := fmt.Sprintf("INSERT INTO %s (user_id,list_id) VALUES ($1,$2)", config.UsersListsTable)
	if _, err := tx.Exec(createUserListQuery, userId, id); err != nil {
		tx.Rollback()
		return 0, err
	}
	return id, tx.Commit()
}
func (t *TodoListPostgres) GetAll(userId int) ([]todo.TodoList, error) {
	var lists []todo.TodoList
	query := fmt.Sprintf("SELECT t1.id, t1.title,t1.description FROM %s t1 INNER JOIN %s t2 on t1.id=t2.list_id WHERE t2.user_id=$1", config.TodoListsTable, config.UsersListsTable)
	err := t.db.Select(&lists, query, userId)
	return lists, err
}

// GetById - Get id list by id user
func (t *TodoListPostgres) GetById(userId int, listId int) (todo.TodoList, error) {
	var list todo.TodoList
	query := fmt.Sprintf(`SELECT t1.id, t1.title,t1.description FROM %s t1 inner join %s t2
		on t1.id=t2.list_id where t1.id=$1 and t2.user_id=$2`, config.TodoListsTable, config.UsersListsTable)
	err := t.db.Get(&list, query, listId, userId)
	return list, err
}

func (t *TodoListPostgres) Delete(userId int, listId int) error {
	query := fmt.Sprintf(`DELETE FROM %s as t1 USING %s as t2 WHERE t2.list_id=t1.id and t1.id=$1 and t2.user_id=$2`, config.TodoListsTable, config.UsersListsTable)
	_, err := t.db.Exec(query, listId, userId)
	return err
}
func (t *TodoListPostgres) Update(userId int, listId int, input todo.TodoListInput) error {
	setValues := make([]string, 0)
	// slice for placeholders
	//args := make([]interface{}, 0)
	//argId := 1
	if input.Title != nil {
		setValues = append(setValues, fmt.Sprintf("title='%s'", *input.Title))

	}
	if input.Description != nil {
		setValues = append(setValues, fmt.Sprintf("description='%s'", *input.Description))
	}
	//strings.Join(setValues, ", ")
	query := fmt.Sprintf(`UPDATE %s t1 SET %s FROM %s t2 WHERE t2.list_id=t1.id and t1.id= $1 and t2.user_id= $2`,
		config.TodoListsTable, strings.Join(setValues, ", "), config.UsersListsTable)
	logrus.Debugf("update Query: %s", query)
	logrus.Debugf("setValues: %s", setValues)
	_, err := t.db.Exec(query, listId, userId)
	return err
}
