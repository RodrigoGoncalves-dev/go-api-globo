package route

import (
	"example.com/go-auth-globo/internal/database"
	"example.com/go-auth-globo/internal/handler"
	"example.com/go-auth-globo/internal/middleware"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func ConfigRoute(
	router *gin.Engine,
	authH *handler.AuthHandler,
	usersH *handler.UserHandler,
) *gin.Engine {

	store, _ := database.OpenRedis()
	router.Use(sessions.Sessions("cookie_auth_globo", store))

	main := router.Group("/api/v1")
	{
		auth := main.Group("/auth")
		{
			auth.POST("/login", authH.DoLogin)
		}
		users := main.Group("/users", middleware.AuthMiddleware())
		{
			users.GET("/:email", usersH.GetUser)
		}
	}

	return router
}
