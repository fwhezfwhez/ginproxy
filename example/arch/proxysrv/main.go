package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.New()

	leafsrv := ginproxy.Group(r, "http://localhost:8081")


	leafsrv.GET("/leaf")
	leafsrv.GET("/leaf/:id")
	leafsrv.POST("/leaf")
	leafsrv.DELETE("/leaf/:id")
	leafsrv.PATCH("/leaf/:id")
	leafsrv.PUT("/leaf/:id")

	r.Run(":8080")
}