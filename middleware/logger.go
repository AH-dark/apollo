package middleware

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func Logger() gin.HandlerFunc {
	return gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		return fmt.Sprintf("[GIN] %s | %s %3d %s | %12s | %16s | %-36s | %s %-6s %s %s\n",
			param.TimeStamp.Format("2006-01-02 15:04:05"),
			param.StatusCodeColor(), param.StatusCode, param.ResetColor(),
			param.Latency,
			param.ClientIP,
			param.Keys["request_id"],
			param.MethodColor(), param.Method, param.ResetColor(),
			param.Path,
		)
	})
}
