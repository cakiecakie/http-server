package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"git.llsapp.com/cc/pkg/api/gin/middleware/auditlog"
)

func TestReq(c *gin.Context) {

	log.Printf("content_length: %d", c.Request.ContentLength)
	c.JSON(http.StatusOK, gin.H{})
}

func main() {

	engine := gin.New()

	engine.Use(
		auditlog.Auditlog(true, 1024),
	)

	apiGroup := engine.Group("/api")

	apiGroup.POST("/test", TestReq)

	http.ListenAndServe(":8080", engine)
}
