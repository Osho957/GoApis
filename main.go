package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"osho.com/db"
	"osho.com/routes"
)

func main() {
	fmt.Println("Starting the server")
	db.InitDB()
	server := gin.Default()
	routes.RegisterRoutes(server)
	server.Run(":8080") // localhost:8080
}
