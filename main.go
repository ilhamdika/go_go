package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	initDB()
	defer db.Close()

	r := gin.Default()

	r.POST("/users", createUser)
	r.GET("/users", getUsers)
	r.GET("/users/:id", getUser)
	r.PUT("/users/:id", updateUser)
	r.DELETE("/users/:id", deleteUser)

	r.Run(":8080")
}