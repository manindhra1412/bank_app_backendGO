package api

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/manindhra1412/simple_bank/token"
)

const (
	authorizationHeaderKey  = "authorization"
	authorizationTypeBearer = "bearer"
	authorizationPayloadKey = "authorization_payload"
)

func authMiddleware(tokenMaker token.Maker) gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.GetHeader("Authorization") == "" {
			err := fmt.Errorf("Authorization header is not provided")
			c.JSON(http.StatusUnauthorized, errorResponse(err))
			c.Abort()
			return
		}

		header := strings.Split(c.GetHeader("Authorization"), "Bearer ")
		if len(header) < 2 {
			err := fmt.Errorf("Authorization header is not valid")
			c.JSON(http.StatusUnauthorized, errorResponse(err))
			c.Abort()
			return
		}

		token := header[1]
		payload, err := tokenMaker.VerifyToken(token)
		if err != nil {
			c.JSON(http.StatusUnauthorized, errorResponse(err))
			c.Abort()
			return
		}

		c.Set(authorizationPayloadKey, payload)
		c.Next()
	}
}
