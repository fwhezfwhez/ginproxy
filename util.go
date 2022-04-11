package util

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"time"
)

var cli = http.Client{
	Timeout: 15 * time.Second,
}

func handleParam(urlptr *string, c *gin.Context) {
	if len(c.Params) > 0 {
		for _, v := range c.Params {
			*urlptr = strings.Replace(*urlptr, fmt.Sprintf(":%s", v.Key), v.Value, -1)
		}
	}
}

func handleQuery(urlptr *string, c *gin.Context) {
	*urlptr = fmt.Sprintf("%s?%s", *urlptr, c.Request.URL.RawQuery)
}
