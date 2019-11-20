package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type Person struct {
	Name    string `form:"name"`
	Address string `form:"address"`
}

func main() {
	r := gin.Default()

	r.Any("/testing", startPage)
	r.Run(":8080")
}

func startPage(ctx *gin.Context) {
	var person Person
	if ctx.ShouldBindQuery(&person) == nil {
		log.Println("===Only Bind By Query String===")
		log.Println(person.Name)
		log.Println(person.Address)
	}

	ctx.String(http.StatusOK, "Success")
}
