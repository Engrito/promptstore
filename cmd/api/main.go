package main

import (
	"log"

	"github.com/engrito/promptstore/internal/user"
	"github.com/gin-gonic/gin"
)

func routes(userHanlders user.UserHandlers) *gin.Engine {
	r := gin.Default()
	r.POST("/api/v1/user/signup", userHanlders.Singup)
	r.GET("/api/v1/user/login", userHanlders.Login)
	r.GET("/api/v1/user", userHanlders.GetById)
	r.GET("/api/v1/user/users", userHanlders.GetAll)
	r.PATCH("/api/v1/user/", userHanlders.UpdateByID)
	r.DELETE("/api/v1/user", userHanlders.DeleteByID)
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
