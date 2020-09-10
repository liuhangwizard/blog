package middleware

import (
	"github.com/gin-gonic/gin"
	"time"
	VBLog "vb-server/utils/log"
)

func VBMiddlewareLog(c *gin.Context) {

	//record start
	reqStart := time.Now()

	// request
	c.Next()

	//record end

	method := c.Request.Method
	url := c.Request.URL.String()
	funcName := c.HandlerName()
	reqEnd := time.Since(reqStart).String()

	//log content
	content := "[" + method + "] " + reqEnd + " " + url + " [" + funcName + "]"

	VBLog.Print.Info(content)
}
