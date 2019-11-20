package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

//  curl  -H "Content-Type: application/x-www-form-urlencoded" -d  'names[first]=thinkerou&names[second]=tianou' http://127.0.0.1:8080/post\?ids\[a\]\=1234\&ids\[b\]\=hello
func main() {
	router := gin.Default()

	router.POST("/post", func(c *gin.Context) {

		ids := c.QueryMap("ids")
		names := c.PostFormMap("names")

		fmt.Printf("ids: %v; names: %v", ids, names)
	})
	router.Run(":8080")
}
