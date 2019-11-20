package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

// curl -d 'name=jackie&message=msg' -H "Content-type: application/x-www-form-urlencoded"  http://127.0.0.1:8080/POST\?id\=122\&page\=2222

func main() {
	router := gin.Default()

	router.POST("/POST", func(context *gin.Context) {
		id := context.Query("id")
		page := context.DefaultQuery("page", "0")
		name := context.PostForm("name")
		message := context.PostForm("message")

		fmt.Println("id: %s; page: %s; name: %s; message: %s", id, page, name, message)
	})

	router.Run(":8080")
}
