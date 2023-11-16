package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Message struct {
	Msg string `json:"text"`
	ID int 
}

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World",
		})
	})

	r.POST("/post", func(c *gin.Context) {
		var message Message
		if err := c.BindJSON(&message); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"received": message.Msg})
	})

	r.Run(":3000") 
}