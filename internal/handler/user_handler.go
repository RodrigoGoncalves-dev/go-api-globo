package handler

import (
	"net/http"
	"time"

	"example.com/go-auth-globo/internal/usecase"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	uc *usecase.UserUsecase
}

func (h *UserHandler) GetUser(c *gin.Context) {
	email := c.Param("email")

	u, err := h.uc.GetUser(email)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Ocorreu um erro, entre em contato com o time de suporte."})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"user":  u,
		"endAt": time.Now().Format("2006-01-02T15:04:05.000Z"),
	})
}

func NewUserHandler(usecase *usecase.UserUsecase) *UserHandler {
	return &UserHandler{
		usecase,
	}
}
