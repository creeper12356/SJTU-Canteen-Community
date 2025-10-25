package middleware

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		userID := session.Get("user_id")
		if userID == nil {
			c.JSON(401, gin.H{"error": "Not logged in"})
			c.Abort()
			return
		}

		c.Set("user_id", userID.(uint))
		c.Next()
	}
}
