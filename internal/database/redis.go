package database

import (
	"example.com/go-auth-globo/internal/config"
	"example.com/go-auth-globo/internal/service"
	redisStore "github.com/gin-contrib/sessions/redis"
)

func OpenRedis() (redisStore.Store, error) {
	url := config.AppInfo.REDIS_URL
	user := config.AppInfo.REDIS_USER
	password := config.AppInfo.REDIS_PASSWORD
	secret := config.AppInfo.REDIS_SECRET

	store, err := redisStore.NewStore(10, "tcp", url, user, password, []byte(secret))

	if err != nil {
		service.Logger().Error("Error on parse URL")
		panic(err)
	}

	return store, err
}
