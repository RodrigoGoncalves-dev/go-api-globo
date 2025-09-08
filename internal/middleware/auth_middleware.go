package middleware

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		token := session.Get("token")

		if token == nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Não autorizado"})
			return
		}

		c.Next()
	}
}
