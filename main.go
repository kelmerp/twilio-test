package main

import (
	"net/http"

	"github.com/davecgh/go-spew/spew"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.POST("/webhook", func(c *gin.Context) {
		var webhookPostBody = make(map[string]interface{})
		c.ShouldBindJSON(&webhookPostBody)
		spew.Dump(webhookPostBody)
		c.JSON(http.StatusOK, gin.H{"status": "you are logged in"})
	})

	r.Run() // run on default port :8080
}
