package main

import (
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	router := gin.Default()

	router.GET("/someGet", middleware1, middleware2, handler)

	router.Run()
}

func handler(context *gin.Context) {
	log.Println("exec handler")
}

func middleware1(context *gin.Context) {
	log.Println("arrive at middleware1")

	context.Next()

	log.Println("exec middleware1")
}

func middleware2(context *gin.Context) {
	log.Println("arrive at middleware2")

	context.Next()

	log.Println("exec middleware2")
}
