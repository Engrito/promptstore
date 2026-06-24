package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// define simple get endpoint
	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"message": "pong"})
	})

	log.Println("server, http://localhost:8080")
	err := r.Run()
	log.Println("failed to start server :", err)
}
