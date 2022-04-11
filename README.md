## ginproxy
ginproxy is a toolkit which helps build  a proxy application in golang using github.com/gin-gonic/gin.

Which case you can use ginproxy:

- You determine to put a stateless proxy in front of your real service leaf application.

- You want migrate a real service application to a stateless proxy application.(This is the real reason to design it.)

## Start

```go
go get github.com/fwhezfwhez/ginproxy
```

```go
var xxLeafsrv *ginproxy.NewGroup
func init() {
    var proxyUrl string
	switch config.Mode {
	case "local":
		proxyUrl = "http://<local_public_ip_port>/leafsrv"
	case "dev":
		proxyUrl = "http://<dev_inner_ip_port>/leafsrv"
	case "pro":
		proxyUrl = "http://xxx.inner-domain.com/leafsrv"
	}
	xxLeafsrv = util.Group(r, proxyUrl)
}

func main(){
    r := gin.New()

    g := r.Group("/api")
    // old
    /*
        r.GET("/leaf", auth, getLeafs)
        r.GET("/leaf/:id", auth, getLeaf)
        r.POST("/leaf", auth, addLeaf)
        r.DELETE("/leaf/:id", auth, deleteLeaf)
        r.PATCH("/leaf/:id", auth, updateLeafField)
        r.PUT("/leaf/:id", auth, updateLeafAllFields)
    */

    // ginproxy

    leafsrv = ginproxy.Group(g, proxyUrl)

    leafsrv.GET("/leaf", auth, getLeafs)
    leafsrv.GET("/leaf/:id", auth, getLeaf)
    leafsrv.POST("/leaf", auth, addLeaf)
    leafsrv.DELETE("/leaf/:id", auth, deleteLeaf)
    leafsrv.PATCH("/leaf/:id", auth, updateLeafField)
    leafsrv.PUT("/leaf/:id", auth, updateLeafAllFields)


    r.Run(":8080")

}

```