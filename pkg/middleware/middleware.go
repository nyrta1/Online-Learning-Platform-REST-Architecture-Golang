package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"online-learning-platform/internal/models"
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

	logger.Info("User is authenticated!")
	c.Next()
}

func RequireTeacherMiddleware(c *gin.Context) {
	logger := logger.GetLogger()

	userRole, exists := c.Get("roles")
	if !exists {
		logger.Error("Roles not found in context")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Roles not found"})
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	isTeacher := false
	if userRole == "TEACHER" {
		isTeacher = true
	}

	if !isTeacher {
		logger.Error("User is not a teacher")
		c.JSON(http.StatusForbidden, gin.H{"error": "User is not a teacher"})
		c.AbortWithStatus(http.StatusForbidden)
		return
	}

	logger.Info("User role is okay!")
	c.Next()
}

func RequireAdminMiddleware(c *gin.Context) {
	logger := logger.GetLogger()

	userRoles, exists := c.Get("role")
	if !exists {
		logger.Error("Roles not found in context")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Roles not found"})
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	roles, ok := userRoles.([]models.Role)
	if !ok {
		logger.Error("Invalid role type in context")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid role type"})
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	isAdmin := false
	for _, role := range roles {
		if role.Name == "ADMIN" {
			isAdmin = true
			break
		}
	}

	if !isAdmin {
		logger.Error("User is not a admin")
		c.JSON(http.StatusForbidden, gin.H{"error": "User is not a admin"})
		c.AbortWithStatus(http.StatusForbidden)
		return
	}

	logger.Info("User role is okay!")
	c.Next()
}
