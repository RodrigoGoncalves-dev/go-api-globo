package route

import (
	"time"

	"example.com/go-auth-globo/internal/database"
	"example.com/go-auth-globo/internal/handler"
	"example.com/go-auth-globo/internal/middleware"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func ConfigRoute(
	router *gin.Engine,
	authH *handler.AuthHandler,
	usersH *handler.UserHandler,
) *gin.Engine {

	store, _ := database.OpenRedis()
	config := cors.DefaultConfig()

	config.AllowAllOrigins = true
	config.AllowMethods = []string{"POST", "GET", "PUT", "OPTIONS"}
	config.AllowHeaders = []string{"Origin", "Content-Type", "Authorization"}
	config.MaxAge = 10 * time.Hour
	config.AllowCredentials = true

	router.Use(cors.New(config))
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
