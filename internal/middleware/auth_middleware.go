package middleware

import (
	"awesomeProject1/internal/auth"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authorization header missing or malformed"})
			return
		}

		tokenStr := strings.TrimPrefix(authHeader, "Bearer ")

		_, claims, err := auth.ValidateJWT(tokenStr)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
			return
		}

		userIDFloat, ok := claims["user_id"].(float64)
		if !ok {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid token payload"})
			return
		}

		role, ok := claims["role"].(string)
		if !ok {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid token payload"})
			return
		}

		c.Set("userID", uint(userIDFloat))
		c.Set("userRole", role)
		c.Next()
	}
}

func RequireRole(requiredRole string) gin.HandlerFunc {
	return func(c *gin.Context) {
		role, exists := c.Get("userRole")
		if !exists || role != requiredRole {
			c.JSON(http.StatusForbidden, gin.H{"error": "Доступ запрещён"})
			c.Abort()
			return
		}
		c.Next()
	}
}

func GetUserFromContext(c *gin.Context) (uint, string, bool) {
	userID, exists := c.Get("userID")
	if !exists {
		return 0, "", false
	}

	role, exists := c.Get("userRole")
	if !exists {
		return 0, "", false
	}

	return userID.(uint), role.(string), true
}
