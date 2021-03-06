package main

import (
	"net/http"

	"github.com/fullstacker-go/concurr_gin/handler"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/test/:id", handler.TestHandler)
	r.GET("/test", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello From Gin...")
	})
	r.Run(":3000")
}
