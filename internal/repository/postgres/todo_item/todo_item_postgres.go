package todo_item

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	todo "github.com/katenester/Todo/internal/models"
	"github.com/katenester/Todo/internal/repository/postgres/config"
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
	createItemQuery := fmt.Sprintf("INSERT INTO %s (title,description) VALUES ($1,$2) RETURNING id", config.TodoItemsTable)
	row := tx.QueryRow(createItemQuery, item.Title, item.Description)
	err = row.Scan(&itemId)
	if err != nil {
		tx.Rollback()
		return 0, err
	}
	createListItemQuery := fmt.Sprintf("INSERT INTO %s (item_id,list_id) VALUES ($1,$2)", config.ListsItemsTable)
	_, err = tx.Exec(createListItemQuery, itemId, listId)
	if err != nil {
		tx.Rollback()
		return 0, err
	}
	return itemId, tx.Commit()
}
func (t *TodoItemPostgres) GetAll(userId int, listId int) ([]todo.TodoItem, error) {
	var items []todo.TodoItem
	query := fmt.Sprintf(`SELECT item.id, item.title,item.description,item.done FROM %s item 
                                             INNER JOIN %s listItem ON item.id=listItem.item_id
                                             INNER JOIN %s userList ON userList.list_id=listItem.list_id
                                             WHERE userList.user_id=$1 AND userList.list_id=$2`, config.TodoItemsTable, config.ListsItemsTable, config.UsersListsTable)
	return items, t.db.Select(&items, query, userId, listId)
}
func (t *TodoItemPostgres) GetById(userId int, itemId int) (todo.TodoItem, error) {
	var item todo.TodoItem
	query := fmt.Sprintf(`SELECT items.id,items.title,items.description, items.done FROM %s items 
                                                         INNER JOIN %s listItem ON listItem.item_id=items.id
                                                         INNER JOIN %s usersList ON usersList.list_id=listItem.list_id
                                                         WHERE usersList.user_id=$1 AND items.id=$2`, config.TodoItemsTable, config.ListsItemsTable, config.UsersListsTable)
	return item, t.db.Get(&item, query, userId, itemId)
}
func (t *TodoItemPostgres) Delete(userId int, itemId int) error {
	query := fmt.Sprintf(`DELETE FROM %s item WHERE item.id=$1`, config.TodoItemsTable)
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
	query := fmt.Sprintf(`UPDATE %s items SET %s WHERE items.id=$%d`, config.TodoItemsTable, strings.Join(setValues, ", "), countArgs)
	args = append(args, itemId)
	_, err := t.db.Exec(query, args...)
	return err
}
