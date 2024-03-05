package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"online-learning-platform/pkg/logger"
	"online-learning-platform/pkg/session"
)

func RequireAuthMiddleware(c *gin.Context) {
	logger := logger.GetLogger()

	if !session.IsAuthenticated(c) {
		logger.Error("Session Data not found!")
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	c.Next()

}
