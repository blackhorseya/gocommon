package ginhttp

import (
	"github.com/blackhorseya/gocommon/pkg/contextx"
	"github.com/gin-gonic/gin"
)

// AddContextx add custom contextx middleware
func AddContextx() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("ctx", contextx.Background())

		c.Next()
	}
}
