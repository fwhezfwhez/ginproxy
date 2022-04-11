package util

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"time"
)

type NewGroup struct {
	rg           *gin.RouterGroup
	proxyPassUrl string
	errnoTmpl    errnoTmpl
}

func Group(r gin.IRouter, proxyPassUrl string) (*NewGroup) {
	rg := r.Group("")
	return &NewGroup{
		rg:           rg,
		proxyPassUrl: proxyPassUrl,
		errnoTmpl:    defaultErrnoTmpl,
	}
}


func (ng *NewGroup) httpProxyPass(url string, c *gin.Context) {
	// 处理url内的param参数
	handleParam(&url, c)

	// 处理url内的query参数
	handleQuery(&url, c)

	fmt.Printf("proxy pass to %s %s", c.Request.Method, url)

	req, e := http.NewRequest(c.Request.Method, url, c.Request.Body)
	if e != nil {
		fmt.Printf("proxy pass new request err %s %s %s %v \n",
			time.Now().Format("2006-01-02 15:04:05"), c.FullPath(), c.Request.Method,
			e)

		c.JSON(200, gin.H{
			ng.errnoTmpl.ErrnoFiledName: -1,
			ng.errnoTmpl.ErrmsgFieldName: fmt.Sprintf("%s proxy pass new request err %s %s %v \n",
				time.Now().Format("2006-01-02 15:04:05"), c.FullPath(), c.Request.Method,
				e),
		})

		return
	}

	req.Header = c.Request.Header

	resp, e := cli.Do(req)
	if e != nil {
		fmt.Printf("proxy pass on response err %s %s %s %v \n",
			time.Now().Format("2006-01-02 15:04:05"), c.FullPath(), c.Request.Method,
			e)

		c.JSON(200, gin.H{
			ng.errnoTmpl.ErrnoFiledName: -1,
			ng.errnoTmpl.ErrmsgFieldName: fmt.Sprintf("%s proxy pass on response err %s %s %v \n",
				time.Now().Format("2006-01-02 15:04:05"), c.FullPath(), c.Request.Method,
				e),
		})
		return
	}

	if resp.Body == nil {
		c.JSON(200, gin.H{
			ng.errnoTmpl.ErrnoFiledName: -1,
			ng.errnoTmpl.ErrmsgFieldName: fmt.Sprintf("proxy pass repsonse body nil err %s %s %s \n",
				time.Now().Format("2006-01-02 15:04:05"), c.FullPath(), c.Request.Method,
			),
		})
		return
	}

	for k, v := range resp.Header {
		if len(v) > 0 {
			c.Writer.Header().Set(k, v[0])
		}
	}

	io.Copy(c.Writer, resp.Body)
}


func (ng *NewGroup) SetRespTmpl(errnoFieldName string, errmsgFieldName string) {
	ng.errnoTmpl = newErrnoTmpl(errnoFieldName, errmsgFieldName)
}

func (ng *NewGroup) POST(rel string, handlers ... func(c *gin.Context)) {
	ng.rg.POST(rel, func(c *gin.Context) {
		ng.httpProxyPass(ng.proxyPassUrl+rel, c)
	})
}

func (ng *NewGroup) GET(rel string, handlers ... func(c *gin.Context)) {
	ng.rg.GET(rel, func(c *gin.Context) {
		ng.httpProxyPass(ng.proxyPassUrl+rel, c)
	})
}

func (ng *NewGroup) DELETE(rel string, handlers ... func(c *gin.Context)) {
	ng.rg.DELETE(rel, func(c *gin.Context) {
		ng.httpProxyPass(ng.proxyPassUrl+rel, c)
	})
}
func (ng *NewGroup) PATCH(rel string, handlers ... func(c *gin.Context)) {
	ng.rg.PATCH(rel, func(c *gin.Context) {
		ng.httpProxyPass(ng.proxyPassUrl+rel, c)
	})
}

func ProxyPassDeclare(c *gin.Context) {
	// 本业务已被代理转发到了其他服务
}
