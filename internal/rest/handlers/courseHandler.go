package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"online-learning-platform/internal/models"
	"online-learning-platform/internal/repository"
	"online-learning-platform/internal/rest/forms"
	"online-learning-platform/pkg/logger"
	"online-learning-platform/pkg/session"
	"strconv"
)

type CourseHandler struct {
	CourseRepo repository.CourseRepo
}

func NewCourseHandler(courseRepo repository.CourseRepo) *CourseHandler {
	return &CourseHandler{
		CourseRepo: courseRepo,
	}
}

func (h *CourseHandler) CreateCourse(context *gin.Context) {
	var courseForm forms.CourseForm
	if err := context.BindJSON(&courseForm); err != nil {
		logger.GetLogger().Error("Invalid login request:", err)
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	teacherID, err := session.GetLoggedUserID(context)
	if err != nil {
		logger.GetLogger().Error("Invalid userID in session:", err)
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid userID in session"})
		return
	}

	var course models.Course
	course.Name = courseForm.Name
	course.Description = courseForm.Description
	course.OwnerUserID = teacherID

	if err := h.CourseRepo.CreateCourse(&course); err != nil {
		logger.GetLogger().Error("Failed to create course:", err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	logger.GetLogger().Info("Course registered successfully")
	context.JSON(http.StatusOK, gin.H{"message": "Course registered successfully"})
}

func (h *CourseHandler) GetAllCourses(context *gin.Context) {
	courseList, err := h.CourseRepo.GetAllCourses()
	if err != nil {
		logger.GetLogger().Error("Failed to extract courses from db:", err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to extract courses from db: " + err.Error()})
		return
	}

	logger.GetLogger().Info("Course List Extracted successfully")
	context.JSON(http.StatusOK, gin.H{"data": courseList})
}

func (h *CourseHandler) GetCourseByID(context *gin.Context) {
	idStr := context.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		logger.GetLogger().Error("Failed to convert string param to int:", err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to convert string param to int:" + err.Error()})
		return
	}

	courseData, err := h.CourseRepo.GetCourseByID(uint(id))
	if err != nil {
		logger.GetLogger().Error("Failed to extract courses from db:", err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to extract courses from db: " + err.Error()})
		return
	}

	logger.GetLogger().Info("Course Extracted successfully")
	context.JSON(http.StatusOK, gin.H{"data": courseData})
}

func (h *CourseHandler) UpdateCourse(context *gin.Context) {
	idStr := context.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		logger.GetLogger().Error("Failed to convert string param to int:", err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to convert string param to int:" + err.Error()})
		return
	}

	var courseForm forms.CourseForm
	if err := context.BindJSON(&courseForm); err != nil {
		logger.GetLogger().Error("Invalid login request:", err)
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	if err := h.CourseRepo.UpdateCourseByID(uint(id), courseForm); err != nil {
		logger.GetLogger().Error("Failed to update courses from db:", err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update courses from db: " + err.Error()})
		return
	}

	logger.GetLogger().Info("Course Updated successfully")
	context.JSON(http.StatusOK, gin.H{"data": "Course Updated successfully"})
}

func (h *CourseHandler) DeleteCourse(context *gin.Context) {
	idStr := context.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		logger.GetLogger().Error("Failed to convert string param to int:", err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to convert string param to int:" + err.Error()})
		return
	}

	if err := h.CourseRepo.DeleteCourse(uint(id)); err != nil {
		logger.GetLogger().Error("Failed to delete courses from db:", err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete courses from db: " + err.Error()})
		return
	}

	logger.GetLogger().Info("Course Deleted successfully")
	context.JSON(http.StatusOK, gin.H{"data": "Course Deleted successfully"})
}
