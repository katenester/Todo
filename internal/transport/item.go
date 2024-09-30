package transport

import (
	"github.com/gin-gonic/gin"
	todo "github.com/katenester/Todo/internal/models"
	"net/http"
	"strconv"
)

func (h *Handler) createItem(c *gin.Context) {
	userId, ok := getUserId(c)
	if ok != nil {
		newErrorResponse(c, http.StatusInternalServerError, ok.Error())
		return
	}
	listId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid list id param")
		return
	}
	var input todo.TodoItem
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	id, err := h.service.TodoItem.Create(userId, listId, input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{"id": id})
}
func (h *Handler) getAllItems(c *gin.Context) {
	UserId, ok := getUserId(c)
	if ok != nil {
		newErrorResponse(c, http.StatusInternalServerError, ok.Error())
		return
	}
	ListId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid list id param")
		return
	}
	items, err := h.service.TodoItem.GetAll(UserId, ListId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, items)
}
func (h *Handler) getItemById(c *gin.Context) {
	UserId, ok := getUserId(c)
	if ok != nil {
		newErrorResponse(c, http.StatusInternalServerError, ok.Error())
		return
	}
	itemId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid list id param")
		return
	}
	item, err := h.service.TodoItem.GetById(UserId, itemId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, item)

}
func (h *Handler) updateItem(c *gin.Context) {
	UserId, ok := getUserId(c)
	if ok != nil {
		newErrorResponse(c, http.StatusInternalServerError, ok.Error())
		return
	}
	itemId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid list id param")
		return
	}
	var input todo.TodoItemInput
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	err = h.service.TodoItem.Update(UserId, itemId, input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, statusResponse{
		Status: "success",
	})
}
func (h *Handler) deleteItem(c *gin.Context) {
	UserId, ok := getUserId(c)
	if ok != nil {
		newErrorResponse(c, http.StatusInternalServerError, ok.Error())
		return
	}
	itemId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid list id param")
		return
	}
	err = h.service.TodoItem.Delete(UserId, itemId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, statusResponse{
		Status: "success",
	})
}
