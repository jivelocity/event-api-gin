package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jivelocity/db"
	"github.com/jivelocity/routes"
)

func main() {
	db.InitDB()
	server := gin.Default()

	routes.RegisterRoutes(server)

	server.Run(":8080")
}
