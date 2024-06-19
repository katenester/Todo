package handler

import (
	"github.com/gin-gonic/gin"
	todo "github.com/katenester/Todo"
	"github.com/sirupsen/logrus"
	"net/http"
)

// Регистрация
func (h *Handler) sigUp(c *gin.Context) {
	var input todo.User
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
	}
	id, err := h.service.Authorization.CreateUser(input)
	logrus.Info(id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

type signInInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// Аутентификация
func (h *Handler) sigIp(c *gin.Context) {
	var input signInInput
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
	}
	token, err := h.service.Authorization.GenerateToken(input.Username, input.Password)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"token": token,
	})

}
