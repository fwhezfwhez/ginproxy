package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.New()

	r.GET("/leaf", auth, getLeafs)
	r.GET("/leaf/:id", auth, getLeaf)
	r.POST("/leaf", auth, addLeaf)
	r.DELETE("/leaf/:id", auth, deleteLeaf)
	r.PATCH("/leaf/:id", auth, updateLeafField)
	r.PUT("/leaf/:id", auth, updateLeafAllFields)

	r.Run(":8081")
}

func auth(c *gin.Context) {
	fmt.Println("auth pass")
	c.Next()
}

func getLeafs(c *gin.Context) {
	fmt.Println("get leafs")
	c.JSON(200, gin.H{
		"errno": 0,
		"data":  []gin.H{},
	})
}

func getLeaf(c *gin.Context) {
	fmt.Println("get a leaf")
	c.JSON(200, gin.H{
		"errno": 0,
		"data":  gin.H{},
	})
}

func addLeaf(c *gin.Context) {
	fmt.Println("add a leaf")
	c.JSON(200, gin.H{
		"errno": 0,
		"data":  gin.H{},
	})
}

func deleteLeaf(c *gin.Context) {
	fmt.Println("delete a leaf")
	c.JSON(200, gin.H{
		"errno": 0,
	})
}

func updateLeafAllFields(c *gin.Context) {
	fmt.Println("update leaf all fields")
	c.JSON(200, gin.H{
		"errno": 0,
	})
}

func updateLeafField(c *gin.Context) {
	fmt.Println("update leaf some fields")
	c.JSON(200, gin.H{
		"errno": 0,
	})
}
