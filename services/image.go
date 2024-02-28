package services

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func ImageUproad(c *gin.Context) {
	file, err := c.FormFile("image")

	if err != nil {
		panic(err)
		return
	}

	err = c.SaveUploadedFile(file, "images/uploads/"+file.Filename)

	if err != nil {
		panic(err)
		return
	}

	c.HTML(http.StatusOK, "index.html", gin.H{
		"image": "images/uploads/" + file.Filename,
	})
}
