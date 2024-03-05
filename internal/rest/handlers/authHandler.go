package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"online-learning-platform/internal/models"
	"online-learning-platform/internal/repository"
	"online-learning-platform/internal/rest/forms"
	"online-learning-platform/pkg/hashing"
	"online-learning-platform/pkg/logger"
	"online-learning-platform/pkg/session"
)

type AuthHandlers struct {
	UserRepo repository.UserRepo
	RoleRepo repository.RoleRepo
}

func NewAuthHandlers(userRepo repository.UserRepo, roleRepo repository.RoleRepo) *AuthHandlers {
	return &AuthHandlers{
		UserRepo: userRepo,
		RoleRepo: roleRepo,
	}
}

func (h *AuthHandlers) WhoAmI(context *gin.Context) {
	loggedUserID, err := session.GetLoggedUserID(context)
	if err != nil {
		logger.GetLogger().Error("Failed to get session user id:", err)
		context.JSON(http.StatusBadRequest, gin.H{"error": "Failed to get session user id"})
		return
	}

	user, err := h.UserRepo.GetUserByID(uint(loggedUserID))
	if err != nil {
		logger.GetLogger().Error("Failed to get user:", err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	logger.GetLogger().Info("User data fetched successfully")
	context.JSON(http.StatusOK, gin.H{"message": "User data fetched successfully", "data": user})
}

func (h *AuthHandlers) Register(context *gin.Context) {
	logger.GetLogger().Info("Received registration request")

	var user models.User

	if err := context.BindJSON(&user); err != nil {
		logger.GetLogger().Error("Invalid registration request:", err)
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	_, err := h.UserRepo.GetUserByUsername(user.Username)
	if err == nil {
		logger.GetLogger().Error("Account already registered for username:", user.Username)
		context.JSON(http.StatusInternalServerError, gin.H{"error": "The account is already registered"})
		return
	}

	role, err := h.RoleRepo.GetByName("USER")
	if err != nil {
		logger.GetLogger().Error("Account cannot get userTypeID:", err.Error())
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Account cannot get userTypeID"})
		return
	}
	user.Roles = append(user.Roles, *role)

	hashedPassword, err := hashing.HashPassword(user.Password)
	if err != nil {
		logger.GetLogger().Error("Unable to hash the password")
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to hash the password"})
		return
	}
	user.Password = hashedPassword

	if err := h.UserRepo.CreateUser(&user); err != nil {
		logger.GetLogger().Error("Failed to create user:", err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	logger.GetLogger().Info("User registered successfully")
	context.JSON(http.StatusOK, gin.H{"message": "User registered successfully"})
}

func (h *AuthHandlers) Login(context *gin.Context) {
	var data forms.LoginForm
	if err := context.BindJSON(&data); err != nil {
		logger.GetLogger().Error("Invalid login request:", err)
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	user, err := h.UserRepo.GetUserByUsername(data.Username)
	if err != nil {
		logger.GetLogger().Error("Failed to get user by username:", err)
		context.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	if !hashing.CheckPasswordHash(data.Password, user.Password) {
		logger.GetLogger().Error("Authentication failed for username:", user.Username)
		context.JSON(http.StatusUnauthorized, gin.H{"error": "Authentication failed"})
		return
	}

	if err := session.SetSessionLoggedID(context, user.ID); err != nil {
		logger.GetLogger().Error("Can't authenticate user. The issue is:", err.Error())
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Can't authenticate user"})
		return
	}

	logger.GetLogger().Info("User login successful")
	context.JSON(http.StatusOK, gin.H{"message": "User login successful"})
}

func (h *AuthHandlers) Logout(context *gin.Context) {
	logger.GetLogger().Info("User logout")

	if err := session.RemoveSessionUserID(context); err != nil {
		logger.GetLogger().Error("Can't remove the authenticated user. The issue is:", err.Error())
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Can't remove the authenticated user"})
		return
	}

	logger.GetLogger().Info("User logged out successfully")
	context.JSON(http.StatusOK, gin.H{"status": "success", "message": "User logged out successfully", "data": nil})
}

func (h *AuthHandlers) Update(context *gin.Context) {
	var data forms.UpdateForm
	if err := context.BindJSON(&data); err != nil {
		logger.GetLogger().Error("Invalid login request:", err)
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	loggerUserID, err := session.GetLoggedUserID(context)
	if err != nil {
		logger.GetLogger().Error("Failed to get user id:", err)
		context.JSON(http.StatusBadRequest, gin.H{"error": "Failed to get user id"})
		return
	}

	if err := h.UserRepo.UpdateUser(uint(loggerUserID), data); err != nil {
		logger.GetLogger().Error("Failed to update user:", err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	logger.GetLogger().Info("User updated successfully")
	context.JSON(http.StatusOK, gin.H{"status": "success", "message": "User updated successfully"})
}

func (h *AuthHandlers) Delete(context *gin.Context) {
	loggerUserID, err := session.GetLoggedUserID(context)
	if err != nil {
		logger.GetLogger().Error("Failed to get user id:", err)
		context.JSON(http.StatusBadRequest, gin.H{"error": "Failed to get user id"})
		return
	}

	if err := h.UserRepo.DeleteUser(uint(loggerUserID)); err != nil {
		logger.GetLogger().Error("Failed to delete user:", err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	logger.GetLogger().Info("User deleted successfully")
	context.JSON(http.StatusOK, gin.H{"status": "success", "message": "User deleted successfully"})
}
