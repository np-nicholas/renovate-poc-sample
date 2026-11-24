package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	_ "golang.org/x/net/html"
)

func main() {
	r := gin.Default()
	r.GET("/healthz", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})
	_ = r.Run(":8080")
}
