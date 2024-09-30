package transport

import (
	"github.com/gin-gonic/gin"
	todo "github.com/katenester/Todo/internal/models"
	"net/http"
	"strconv"
)

// @Summary Create todo item
// @Security ApiKeyAuth
// @Tags item
// @Description create todo item
// @ID create-item
// @Accept  json
// @Produce  json
// @Param id path int true "List ID"
// @Param input body models.TodoItem true "item info"
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/lists/:id/items [post]
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

// @Summary GetAll todo item
// @Security ApiKeyAuth
// @Tags item
// @Description GetAll todo item
// @ID getAll-item
// @Accept  json
// @Produce  json
// @Param id path int true "List ID"
// @Success 200 {integer}  []models.TodoItem
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/lists/:id/items [get]
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

// @Summary Get by id todo item
// @Security ApiKeyAuth
// @Tags item
// @Description Get by id todo item
// @ID get-by-id-item
// @Accept  json
// @Produce  json
// @Param id path int true "Item ID"
// @Success 200 {integer} models.TodoItem
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/items [get]
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

// @Summary Update todo item
// @Security ApiKeyAuth
// @Tags item
// @Description Update id todo item
// @ID update-item
// @Accept  json
// @Produce  json
// @Param id path int true "Item ID"
// @Param input body models.TodoItem true "item info"
// @Success 200 {integer} statusResponse
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/items [put]
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

// @Summary Delete todo item
// @Security ApiKeyAuth
// @Tags item
// @Description Delete id todo item
// @ID delete-item
// @Accept  json
// @Produce  json
// @Param id path int true "Item ID"
// @Success 200 {integer} statusResponse
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/items [delete]
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
