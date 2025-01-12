package middleware

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gabrielfmcoelho/platform-core/domain"
	"github.com/gabrielfmcoelho/platform-core/internal/tokenutil"
	"github.com/gin-gonic/gin"
)

func JwtAuthMiddleware(secret string) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")
		// console log the authHeader
		fmt.Println("Authorization header")
		t := strings.Split(authHeader, " ")
		if len(t) == 2 {
			authToken := t[1]
			authorized, err := tokenutil.IsAuthorized(authToken, secret)
			skip, _ := tokenutil.SkipTokenValidation(authToken)
			if skip {
				c.Next()
				fmt.Println("Skip token validation")
				return
			}
			if authorized {
				userID, err := tokenutil.ExtractIDFromToken(authToken, secret)
				if err != nil {
					c.JSON(http.StatusUnauthorized, domain.ErrorResponse{Message: err.Error()})
					c.Abort()
					fmt.Println("Error extracting ID from token")
					return
				}
				c.Set("x-user-id", userID)
				c.Next()
				fmt.Println("Authorized")
				return
			}
			c.JSON(http.StatusUnauthorized, domain.ErrorResponse{Message: err.Error()})
			c.Abort()
			fmt.Println("Not authorized")
			return
		}
		c.JSON(http.StatusUnauthorized, domain.ErrorResponse{Message: "Not authorized"})
		c.Abort()
		fmt.Println("Not authorized")
	}
}
