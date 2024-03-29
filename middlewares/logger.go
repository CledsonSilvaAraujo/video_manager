package middlewares

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

func Logger() gin.HandlerFunc {
	return gin.LoggerWithFormatter(func(params gin.LogFormatterParams) string {
		return fmt.Sprintf("%s - [%s] \"%s %s %s %s %s\"\n",
			params.TimeStamp.Format(time.RFC1123),
			params.Method,
			params.Path,
			fmt.Sprintf("%d", params.StatusCode),
			params.Latency,
			params.ClientIP,
			params.ErrorMessage)
	})
}
