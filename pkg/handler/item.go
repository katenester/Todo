package handler

import (
	"github.com/gin-gonic/gin"
	todo "github.com/katenester/Todo"
	"net/http"
	"strconv"
)

func (h *Handler) createItem(c *gin.Context) {
	UserId, ok := getUserId(c)
	if ok != nil {
		return
	}
	ListId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid list id param")
		return
	}
	var input todo.TodoItem
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
	}
	id, err := h.service.TodoItem.Create(UserId, ListId, input)
}
func (h *Handler) getAllItems(c *gin.Context) {

}
func (h *Handler) getItemById(c *gin.Context) {

}
func (h *Handler) updateItem(c *gin.Context) {

}
func (h *Handler) deleteItem(c *gin.Context) {

}
