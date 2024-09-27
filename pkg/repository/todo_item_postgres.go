package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	todo "github.com/katenester/Todo"
	"strings"
)

type TodoItemPostgres struct {
	db *sqlx.DB
}

func NewTodoItemPostgres(db *sqlx.DB) *TodoItemPostgres {
	return &TodoItemPostgres{db: db}
}

func (t *TodoItemPostgres) Create(listId int, item todo.TodoItem) (int, error) {
	tx, err := t.db.Begin()
	if err != nil {
		return 0, err
	}
	var itemId int
	createItemQuery := fmt.Sprintf("INSERT INTO %s (title,description) VALUES ($1,$2) RETURNING id", todoItemsTable)
	row := tx.QueryRow(createItemQuery, item.Title, item.Description)
	err = row.Scan(&itemId)
	if err != nil {
		tx.Rollback()
		return 0, err
	}
	createListItemQuery := fmt.Sprintf("INSERT INTO %s (item_id,list_id) VALUES ($1,$2)", listsItemsTable)
	_, err = tx.Exec(createListItemQuery, itemId, listId)
	if err != nil {
		tx.Rollback()
		return 0, err
	}
	return itemId, tx.Commit()
}
func (t *TodoItemPostgres) GetAll(userId int, listId int) ([]todo.TodoItem, error) {
	var items []todo.TodoItem
	query := fmt.Sprintf(`SELECT item.id, item.title,item.description,item.done from %s item 
                                             inner join %s listItem on item.id=listItem.item_id
                                             inner join %s userList on userList.list_id=listItem.list_id
                                             where userList.user_id=$1 and userList.list_id=$2`, todoItemsTable, listsItemsTable, usersListsTable)
	return items, t.db.Select(&items, query, userId, listId)
}
func (t *TodoItemPostgres) GetById(userId int, itemId int) (todo.TodoItem, error) {
	var item todo.TodoItem
	query := fmt.Sprintf(`SELECT items.id,items.title,items.description, items.done from %s items inner join %s listItem on listItem.item_id=items.id
                                  inner join %s usersList on usersList.list_id=listItem.list_id
                                  where usersList.user_id=$1 and items.id=$2`, todoItemsTable, listsItemsTable, usersListsTable)
	return item, t.db.Get(&item, query, userId, itemId)
}
func (t *TodoItemPostgres) Delete(userId int, itemId int) error {
	query := fmt.Sprintf(`DELETE FROM %s item where item.id=$1`, todoItemsTable)
	_, err := t.db.Exec(query, itemId)
	return err
}
func (t *TodoItemPostgres) Update(userId int, itemId int, item todo.TodoItemInput) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	countArgs := 1
	if item.Title != nil {
		setValues = append(setValues, fmt.Sprintf("title=$%d", countArgs))
		args = append(args, *item.Title)
		countArgs++
	}
	if item.Description != nil {
		setValues = append(setValues, fmt.Sprintf("description=$%d", countArgs))
		args = append(args, *item.Description)
		countArgs++
	}
	if item.Done != nil {
		setValues = append(setValues, fmt.Sprintf("done=$%d", countArgs))
		args = append(args, *item.Done)
		countArgs++
	}
	query := fmt.Sprintf(`UPDATE %s items set %s where items.id=$%d`, todoItemsTable, strings.Join(setValues, ", "), countArgs)
	args = append(args, itemId)
	_, err := t.db.Exec(query, args...)
	return err
}
