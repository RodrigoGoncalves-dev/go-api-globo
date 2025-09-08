package config

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/viper"
)

type ApplicationInfo struct {
	SECRET_KEY string
	ISSUER     string

	REDIS_URL      string
	REDIS_USER     string
	REDIS_PASSWORD string
	REDIS_SECRET   string

	USER     string
	PASSWORD string
	HOST     string
	PORT     string
	DBNAME   string
}

var AppInfo *ApplicationInfo

func InitAppInfo() {
	AppInfo = &ApplicationInfo{
		SECRET_KEY:     os.Getenv("SECRET_KEY"),
		ISSUER:         os.Getenv("ISSUER"),
		REDIS_URL:      os.Getenv("REDIS_URL"),
		REDIS_USER:     os.Getenv("REDIS_USER"),
		REDIS_PASSWORD: os.Getenv("REDIS_PASSWORD"),
		REDIS_SECRET:   os.Getenv("REDIS_SECRET"),
		USER:           os.Getenv("DB_USER"),
		PASSWORD:       os.Getenv("DB_PASSWORD"),
		HOST:           os.Getenv("DB_HOST"),
		PORT:           os.Getenv("DB_PORT"),
		DBNAME:         os.Getenv("DB_NAME"),
	}
}

func CleanAppInfo() {
	AppInfo = nil
}

func NewViper() *viper.Viper {
	v := viper.New()
	v.AddConfigPath("../configs")
	v.SetConfigType("env")
	env := strings.ToLower(os.Getenv("GOENV"))
	v.SetConfigName(fmt.Sprintf(".%v", env))
	if err := v.ReadInConfig(); err != nil {
		panic(err)
	}
	return v
}
