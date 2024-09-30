package service

import (
	todo "github.com/katenester/Todo/internal/models"
	"github.com/katenester/Todo/internal/repository"
)

type TodoItemService struct {
	repo     repository.TodoItem
	listRepo repository.TodoList
}

func NewTodoItemService(repo repository.TodoItem, listRepo repository.TodoList) *TodoItemService {
	return &TodoItemService{repo: repo, listRepo: listRepo}
}

func (t *TodoItemService) Create(userId, listId int, item todo.TodoItem) (int, error) {
	// Check have list exist or list not belongs to user
	_, err := t.listRepo.GetById(userId, listId)
	if err != nil {
		return 0, err
	}
	return t.repo.Create(listId, item)
}
func (t *TodoItemService) GetAll(userId int, listId int) ([]todo.TodoItem, error) {
	return t.repo.GetAll(userId, listId)
}
func (t *TodoItemService) GetById(userId int, itemId int) (todo.TodoItem, error) {
	return t.repo.GetById(userId, itemId)
}
func (t *TodoItemService) Delete(userId int, itemId int) error {
	return t.repo.Delete(userId, itemId)
}
func (t *TodoItemService) Update(userId int, itemId int, item todo.TodoItemInput) error {
	if err := item.Valid(); err != nil {
		return err
	}
	return t.repo.Update(userId, itemId, item)
}
