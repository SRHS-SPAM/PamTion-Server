package controllers

import (
	"PamTion-Server/services"
	"github.com/gin-gonic/gin"
	"github.com/olahol/go-imageupload"
)

var currentImage *imageupload.Image

func NewController(port string) {
	r := gin.Default()
	r.Static("/images", "./images")
	r.LoadHTMLGlob("templates/*")
	r.MaxMultipartMemory = 2 << 20

	r.POST("/upload", func(c *gin.Context) {
		services.ImageUproad(c)
	})
	r.Run(":8080")
}
