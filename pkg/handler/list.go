package handler

import (
	"github.com/gin-gonic/gin"
	todo "github.com/katenester/Todo"
	"net/http"
	"strconv"
)

func (h *Handler) createList(c *gin.Context) {
	//  Take value UserId from context
	UserId, ok := getUserId(c)
	if ok != nil {
		return
	}
	var input todo.TodoList
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	// call service method
	id, err := h.service.TodoList.Create(UserId, input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{"id": id})

}

type getAllListResponse struct {
	Todos []todo.TodoList `json:"data"`
}

func (h *Handler) getAllLists(c *gin.Context) {
	//  Take value UserId from context
	UserId, ok := getUserId(c)
	if ok != nil {
		return
	}
	// call service method
	list, err := h.service.TodoList.GetAll(UserId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, getAllListResponse{list})
}
func (h *Handler) getListById(c *gin.Context) {
	//  Take value UserId from context
	UserId, ok := getUserId(c)
	if ok != nil {
		return
	}
	ListId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
	}
	// call service method
	list, err := h.service.TodoList.GetById(UserId, ListId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, list)
}
func (h *Handler) updateList(c *gin.Context) {

}
func (h *Handler) deleteList(c *gin.Context) {

}
