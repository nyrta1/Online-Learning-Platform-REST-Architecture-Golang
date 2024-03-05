package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"online-learning-platform/internal/models"
	"online-learning-platform/internal/repository"
	"online-learning-platform/internal/rest/forms"
	"online-learning-platform/pkg/logger"
	"strconv"
)

type LessonHandler struct {
	LessonRepo repository.LessonRepo
}

func NewLessonHandler(lessonRepo repository.LessonRepo) *LessonHandler {
	return &LessonHandler{
		LessonRepo: lessonRepo,
	}
}

func (lh *LessonHandler) GetLessonByID(context *gin.Context) {
	idStr := context.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		logger.GetLogger().Error("Failed to convert string param to int:", err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to convert string param to int:" + err.Error()})
		return
	}

	lesson, err := lh.LessonRepo.GetLessonByID(uint(id))
	if err != nil {
		logger.GetLogger().Error("Failed to get lesson from db:", err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get lesson from db: " + err.Error()})
		return
	}

	logger.GetLogger().Info("Course get successfully")
	context.JSON(http.StatusOK, gin.H{"status": "Course get successfully", "data": lesson})
}

func (lh *LessonHandler) GetAllLessons(context *gin.Context) {
	lessonList, err := lh.LessonRepo.GetAllLessons()
	if err != nil {
		logger.GetLogger().Error("Failed to get all lessons from db:", err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get all lessons from db: " + err.Error()})
		return
	}

	logger.GetLogger().Info("Course get successfully")
	context.JSON(http.StatusOK, gin.H{"status": "Course get successfully", "data": lessonList})
}

func (lh *LessonHandler) CreateLesson(context *gin.Context) {
	var lessonForm forms.LessonForm
	if err := context.BindJSON(&lessonForm); err != nil {
		logger.GetLogger().Error("Invalid login request:", err)
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	var lesson models.Lesson
	lesson.Name = lessonForm.Name
	lesson.Description = lessonForm.Description

	if err := lh.LessonRepo.CreateLesson(&lesson); err != nil {
		logger.GetLogger().Error("Failed to create lesson:", err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	logger.GetLogger().Info("Lesson registered successfully")
	context.JSON(http.StatusOK, gin.H{"message": "Lesson registered successfully"})
}

func (lh *LessonHandler) UpdateLesson(context *gin.Context) {
	idStr := context.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		logger.GetLogger().Error("Failed to convert string param to int:", err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to convert string param to int:" + err.Error()})
		return
	}

	var lessonForm forms.LessonForm
	if err := context.BindJSON(&lessonForm); err != nil {
		logger.GetLogger().Error("Invalid login request:", err)
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	if err := lh.LessonRepo.UpdateLesson(uint(id), lessonForm); err != nil {
		logger.GetLogger().Error("Failed to update lesson from db:", err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update lesson from db: " + err.Error()})
		return
	}

	logger.GetLogger().Info("Lesson Updated successfully")
	context.JSON(http.StatusOK, gin.H{"data": "Lesson Updated successfully"})
}

func (lh *LessonHandler) DeleteLessonByID(context *gin.Context) {
	idStr := context.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		logger.GetLogger().Error("Failed to convert string param to int:", err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to convert string param to int:" + err.Error()})
		return
	}

	if err := lh.LessonRepo.DeleteLesson(uint(id)); err != nil {
		logger.GetLogger().Error("Failed to delete lesson from db:", err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete lesson from db: " + err.Error()})
		return
	}

	logger.GetLogger().Info("Lesson Deleted successfully")
	context.JSON(http.StatusOK, gin.H{"data": "Lesson Deleted successfully"})
}
