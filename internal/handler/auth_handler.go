package handler

import (
	"net/http"

	"example.com/go-auth-globo/internal/domain"
	"example.com/go-auth-globo/internal/usecase"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	u *usecase.AuthUsecase
}

func (h *AuthHandler) DoLogin(c *gin.Context) {
	var input domain.LoginInput

	err := c.ShouldBindJSON(&input)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dados inseridos incorretamente"})
		return
	}

	result, err := h.u.DoLogin(input)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Ocorreu um erro, entre em contato com o time de suporte."})
		return
	}

	session := sessions.Default(c)
	session.Set("email", input.Email)
	session.Set("token", result.Token)
	// MaxAge in 1 hour
	session.Options(sessions.Options{Path: "/", MaxAge: 60 * 60})

	if err = session.Save(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ocorreu um erro, entre em contato com o time de suporte."})
	}

	c.JSON(http.StatusOK, result)
}

func NewAuthHandler(u *usecase.AuthUsecase) *AuthHandler {
	return &AuthHandler{u}
}
