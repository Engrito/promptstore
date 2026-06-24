package main

import (
	"log"

	"github.com/engrito/promptstore/internal/user"
	"github.com/gin-gonic/gin"
)

func routes(userHanlders user.UserHandlers) *gin.Engine {
	r := gin.Default()
	return r
}

func main() {
	userRepo := user.NewInMermoryUserRepository()
	userService := user.NewService(userRepo)
	userHandlers := user.NewUserHandlers(*userService)

	r := routes(*userHandlers)

	log.Println("server, http://localhost:8080")
	err := r.Run()
	log.Println("failed to start server :", err)
}
