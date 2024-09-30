package models

import "errors"

// TodoList структура списка дел для каждого пользователя
type TodoList struct {
	Id          int    `json:"id" db:"id"`
	Title       string `json:"title" db:"title" binding:"required"`
	Description string `json:"description" db:"description"`
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
