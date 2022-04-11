## ginproxy
ginproxy is a toolkit which helps build  a proxy application in golang using github.com/gin-gonic/gin.

Which case you can use ginproxy:

- You determine to put a stateless proxy in front of your real service leaf application.

- You want migrate a real service application to a stateless proxy application.(This is the real reason to design it.)

## Start

```go
go get github.com/fwhezfwhez/ginproxy
```

proxysrv

```go
package main

import (
	"github.com/fwhezfwhez/ginproxy"
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

```

leafsrv
```go
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

```