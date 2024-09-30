package models

import "errors"

// TodoItem представляет дело из списка
type TodoItem struct {
	Id          int    `json:"id" db:"id"`
	Title       string `json:"title" db:"title" binding:"required"`
	Description string `json:"description" db:"description" `
	Done        bool   `json:"done" db:"done"`
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
