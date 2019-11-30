package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

func main() {
	r := gin.Default()

	r.Use(Logger())

	r.GET("/test", func(ctx *gin.Context) {
		example := ctx.MustGet("example").(string)
		name := ctx.MustGet("name").(string)
		log.Println(example)
		log.Println(name)

	})

	r.Run(":8080")
}

func Logger() gin.HandlerFunc {
	return func(context *gin.Context) {
		t := time.Now()

		context.Set("example", "12345")
		context.Set("name", "jackie")

		context.Next()

		latency := time.Since(t)
		log.Println(latency)

		status := context.Writer.Status()
		log.Println(status)
	}
}
