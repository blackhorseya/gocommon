package tokenutil

import (
	"strings"

	"github.com/blackhorseya/gocommon/pkg/er"
	"github.com/gin-gonic/gin"
)

type authHeader struct {
	IDToken string `header:"Authorization"`
}

// GetBearerToken get bearer token from authorization header
func GetBearerToken(c *gin.Context) (string, error) {
	h := authHeader{}
	err := c.ShouldBindHeader(&h)
	if err != nil {
		return "", er.ErrMissingToken
	}

	idTokenHeader := strings.Split(h.IDToken, "Bearer ")
	if len(idTokenHeader) < 2 {
		return "", er.ErrAuthHeaderFormat
	}

	return idTokenHeader[1], nil
}
