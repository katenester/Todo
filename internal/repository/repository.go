package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/katenester/Todo/internal/models"
	"github.com/katenester/Todo/internal/repository/postgres/auth"
	"github.com/katenester/Todo/internal/repository/postgres/todo_item"
	"github.com/katenester/Todo/internal/repository/postgres/todo_list"
)

type Authorization interface {
	CreateUser(user models.User) (int, error)
	GetUser(username, password string) (models.User, error)
}

type TodoList interface {
	Create(userId int, list models.TodoList) (int, error)
	GetAll(userId int) ([]models.TodoList, error)
	GetById(userId int, listId int) (models.TodoList, error)
	Delete(userId int, listId int) error
	Update(userId int, listId int, list models.TodoListInput) error
}

type TodoItem interface {
	Create(listId int, item models.TodoItem) (int, error)
	GetAll(userId int, listId int) ([]models.TodoItem, error)
	GetById(userId int, itemId int) (models.TodoItem, error)
	Delete(userId int, itemId int) error
	Update(userId int, itemId int, item models.TodoItemInput) error
}

type Repository struct {
	Authorization
	TodoList
	TodoItem
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: auth.NewAuthPostgres(db),
		TodoList:      todo_list.NewTodoListPostgres(db),
		TodoItem:      todo_item.NewTodoItemPostgres(db),
	}
}
