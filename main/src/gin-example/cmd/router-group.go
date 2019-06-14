package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()

	v1 := router.Group("/v1")
	{
		v1.GET("/login", routerFun1)
		v1.GET("/submit", routerFun2)
		v1.GET("/read", routerFun3)
	}

	v2 := router.Group("v2")
	{
		v2.POST("/login", routerFun1)
		v2.POST("/submit", routerFun2)
		v2.POST("/read", routerFun3)
	}

	router.Run()
}

func routerFun1(context *gin.Context) {
	context.String(http.StatusOK, "routerFun1")
}

func routerFun2(context *gin.Context) {
	context.String(http.StatusOK, "routerFun2")
}

func routerFun3(context *gin.Context) {
	context.String(http.StatusOK, "routerFun3")
}
