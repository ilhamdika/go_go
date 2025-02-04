package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func apiKeyMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		apiKey := os.Getenv("API_KEY")
		key := c.GetHeader("Authorization")

		if key != apiKey {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		c.Next()
	}
}

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	initDB()
	defer db.Close()

	r := gin.Default()

	protected := r.Group("/")
	protected.Use(apiKeyMiddleware())

	protected.POST("/users", createUser)
	protected.GET("/users", getUsers)
	protected.GET("/users/:id", getUser)
	protected.PUT("/users/:id", updateUser)
	protected.DELETE("/users/:id", deleteUser)

	r.Run(":8080")
}