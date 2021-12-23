package middleware

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		// pre log
		method := c.Request.Method
		path := c.Request.URL.Path
		clientIP := c.ClientIP()
		clientUserAgent := c.Request.UserAgent()
		headers := c.Request.Header
		params := c.Params
		reqID, _ := c.Get("RequestId")

		requestEntry := fmt.Sprintf("reqID=%s, clientIp=%s, method=%s, path=%s, params=%s, headers=%s, userAgent=%s",
			reqID, clientIP, method, path, params, headers, clientUserAgent)

		if len(c.Errors) > 0 {
			log.Errorf("%s, %s", c.Errors.ByType(gin.ErrorTypePrivate).String(), requestEntry)
		} else {
			log.Infof("[GIN] Before Request, %s", requestEntry)
		}

		// process request
		start := time.Now()
		c.Next()
		stop := time.Now()

		// post log
		latency := stop.Sub(start)
		statusCode := c.Writer.Status()

		postEntry := fmt.Sprintf("reqID=%s, clientIP=%s, method=%s, path=%s, statusCode=%v, latency=%s", reqID, clientIP, method, path, statusCode, latency)

		if len(c.Errors) > 0 {
			log.Errorf("%s, %s", c.Errors.ByType(gin.ErrorTypePrivate).String(), postEntry)
		} else {
			log.Infof("[GIN] After Request, %s", postEntry)
		}
	}
}
