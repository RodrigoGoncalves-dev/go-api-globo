package main

import (
	"example.com/go-auth-globo/internal/config"
	"example.com/go-auth-globo/internal/server"
)

func main() {
	config.InitAppInfo()
	defer config.CleanAppInfo()

	server := server.NewServer("5000")
	server.Run()
}
