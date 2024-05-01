package middlewares

import (
	"CarSaleAd-Web-Api/config"
	"CarSaleAd-Web-Api/pkg/logging"
	"bytes"
	"io"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

type bodyLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w *bodyLogWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}
func (w *bodyLogWriter) WriteString(s string) (int, error) {
	w.body.WriteString(s)
	return w.ResponseWriter.WriteString(s)
}

func DefaultStructuredLogger(cfg *config.Config) gin.HandlerFunc {
	logger := logging.NewLogger(cfg)
	return structuredLogger(logger)
}

func structuredLogger(logger logging.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		if strings.Contains(c.FullPath(), "swagger") {
			c.Next()
		} else {
			blw := &bodyLogWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
			start := time.Now() //start
			path := c.FullPath()
			raw := c.Request.URL.RawQuery
			//RequestBody
			bodyBytes, _ := io.ReadAll(c.Request.Body)
			c.Request.Body.Close()
			c.Request.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
			//ResponseBody ---need to have struct
			c.Writer = blw

			c.Next()

			//move from context to param
			param := gin.LogFormatterParams{}
			param.TimeStamp = time.Now()               //stop
			param.Latency = param.TimeStamp.Sub(start) // period time
			param.ClientIP = c.ClientIP()
			param.Method = c.Request.Method
			param.StatusCode = c.Writer.Status()
			param.ErrorMessage = c.Errors.ByType(gin.ErrorTypePrivate).String()
			param.BodySize = c.Writer.Size()

			if raw != "" {
				path = path + "?" + raw
			}
			param.Path = path

			//move from param to Keys
			keys := map[logging.ExtraKey]interface{}{}
			keys[logging.Path] = param.Path
			keys[logging.ClientIP] = param.ClientIP
			keys[logging.Method] = param.Method
			keys[logging.Latency] = param.Latency
			keys[logging.StatusCode] = param.StatusCode
			keys[logging.ErrorMessage] = param.ErrorMessage
			keys[logging.BodySize] = param.BodySize
			keys[logging.RequestBody] = string(bodyBytes)
			keys[logging.ResponseBody] = blw.body.String()

			//info part is called
			logger.Info(logging.RequestResponse, logging.Api, "", keys)
		}
	}

}
