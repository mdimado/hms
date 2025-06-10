package auth

import (
	"net/http"
	"strings"
	"hospital-management/pkg/utils"
	"github.com/gin-gonic/gin"
)

func RequireAuth(jwtSecret string) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			utils.ErrorResponse(c, http.StatusUnauthorized, "Authorization header required", nil)
			c.Abort()
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		if tokenString == authHeader {
			utils.ErrorResponse(c, http.StatusUnauthorized, "Invalid authorization format", nil)
			c.Abort()
			return
		}

		claims, err := utils.ValidateToken(tokenString, jwtSecret)
		if err != nil {
			utils.ErrorResponse(c, http.StatusUnauthorized, "Invalid token", err)
			c.Abort()
			return
		}

		c.Set("user", map[string]interface{}{
			"id":   claims.UserID,
			"role": claims.Role,
		})

		c.Next()
	}
}

func RequireRole(role string) gin.HandlerFunc {
	return func(c *gin.Context) {
		userInterface, exists := c.Get("user")
		if !exists {
			utils.ErrorResponse(c, http.StatusUnauthorized, "User not authenticated", nil)
			c.Abort()
			return
		}

		user := userInterface.(map[string]interface{})
		userRole := user["role"].(string)

		if userRole != role {
			utils.ErrorResponse(c, http.StatusForbidden, "Insufficient permissions", nil)
			c.Abort()
			return
		}

		c.Next()
	}
}
