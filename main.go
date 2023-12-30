package main

import (
	"github.com/gin-gonic/gin"
	"github.com/yoshihyoda/golang_api/db"
	"github.com/yoshihyoda/golang_api/routes"
)

func main() {
	db.InitDB()
	server := gin.Default()

	routes.RegisterRoutes(server)

	server.Run(":8080")
}
