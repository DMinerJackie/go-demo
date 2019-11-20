package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.New()

	r.Use(gin.Recovery())

	authorized := r.Group("/")

	authorized.Use(gin.Logger())
	{
		authorized.GET("/login", func(context *gin.Context) {
			context.JSON(http.StatusOK, gin.H{
				"title": "middleware",
			})
		})
	}

	r.Run(":8080")
}

func loginEndpoint() {

}
