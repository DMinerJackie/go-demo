package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main() {
	router := gin.Default()

	router.POST("/upload", func(context *gin.Context) {
		file, _ := context.FormFile("file")
		log.Println(file.Filename)

		context.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))
	})

	router.Run(":8080")
}
