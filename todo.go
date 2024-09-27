package todo

import "errors"

// TodoList структура списка дел для каждого пользователя
type TodoList struct {
	Id          int    `json:"id" db:"id"`
	Title       string `json:"title" db:"title" binding:"required"`
	Description string `json:"description" db:"description"`
}

// UserList связь n:n для пользователя и его списков
type UserList struct {
	Id     int
	UserId int
	ListId int
}

// TodoItem представляет дело из списка
type TodoItem struct {
	Id          int    `json:"id" db:"id"`
	Title       string `json:"title" db:"title" binding:"required"`
	Description string `json:"description" db:"description" `
	Done        bool   `json:"done" db:"done"`
}

// ListItem связь n:n для списка и дел
type ListItem struct {
	Id     int
	ListId int
	ItemId int
}

type TodoListInput struct {
	Title       *string `json:"title"`
	Description *string `json:"description"`
}

func (t TodoListInput) Valid() error {
	if t.Title == nil && t.Description == nil {
		return errors.New("update structure has not values")
	}
	return nil
}

type TodoItemInput struct {
	Title       *string `json:"title"`
	Description *string `json:"description"`
	Done        *bool   `json:"done"`
}

func (t TodoItemInput) Valid() error {
	if t.Title == nil && t.Description == nil && t.Done == nil {
		return errors.New("update structure has not values")
	}
	return nil
}
