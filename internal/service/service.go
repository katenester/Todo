package service

import (
	"github.com/katenester/Todo/internal/models"
	"github.com/katenester/Todo/internal/repository"
)

type Authorization interface {
	CreateUser(user models.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
}

type TodoList interface {
	Create(userId int, list models.TodoList) (int, error)
	GetAll(userId int) ([]models.TodoList, error)
	GetById(userId int, listId int) (models.TodoList, error)
	Delete(userId int, listId int) error
	Update(userId int, listId int, list models.TodoListInput) error
}

type TodoItem interface {
	Create(userId, listId int, item models.TodoItem) (int, error)
	GetAll(userId int, listId int) ([]models.TodoItem, error)
	GetById(userId int, itemId int) (models.TodoItem, error)
	Delete(userId int, itemId int) error
	Update(userId int, itemId int, item models.TodoItemInput) error
}

type Service struct {
	Authorization
	TodoList
	TodoItem
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		TodoList:      NewTodoListService(repos.TodoList),
		TodoItem:      NewTodoItemService(repos.TodoItem, repos.TodoList),
	}
}
