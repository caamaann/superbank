package middleware

import (
	"net/http"
	"strings"

	"superbank/internal/service"
	"superbank/pkg/util"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware(authService service.AuthService) gin.HandlerFunc {
	return func(c *gin.Context) {
		
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			util.NewResponse("Authorization header required", nil, http.StatusUnauthorized).ReturnGin(c)
			c.Abort()
			return
		}

		tokenString := authHeader[7:]
		userID, role, err := authService.ValidateToken(tokenString)
		if err != nil {
			util.NewResponse("Invalid token", nil, http.StatusUnauthorized).ReturnGin(c)
			c.Abort()
			return
		}

		
		c.Set("userID", userID)
		c.Set("role", role)

		c.Next()
	}
}