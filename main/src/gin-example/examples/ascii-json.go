package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func main() {
	r := gin.Default()

	fmt.Println(strconv.Unquote(strconv.QuoteToASCII("zhong中国")))

	r.GET("/someJson", func(context *gin.Context) {
		data := map[string]interface{}{
			"lang": "GO语言",
			"tag":  "<br>",
		}

		context.JSON(http.StatusOK, data)
	})

	r.Run(":8080")
}
