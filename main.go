package main

import (
	"backend/api"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Register API routes
	api.RegisterRoutes(r)

	r.Run(":8080")
}
