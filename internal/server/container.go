package server

import (
	"database/sql"
	"errors"

	"example.com/go-auth-globo/internal/database"
	"example.com/go-auth-globo/internal/handler"
	"example.com/go-auth-globo/internal/repository"
	"example.com/go-auth-globo/internal/usecase"
	"github.com/gin-gonic/gin"
	"go.uber.org/dig"
)

func NewContainer() *dig.Container {
	c := dig.New()

	if err := errors.Join(
		c.Provide(gin.Default),

		c.Provide(func() *sql.DB {
			db, err := database.OpenDB()
			if err != nil {
				panic(err)
			}
			return db
		}),

		c.Provide(repository.NewAuthRepository),
		c.Provide(usecase.NewAuthUsecase),
		c.Provide(handler.NewAuthHandler),

		c.Provide(repository.NewUserRepository),
		c.Provide(usecase.NewUserUsecase),
		c.Provide(handler.NewUserHandler),
	); err != nil {
		panic(err)
	}

	return c
}
