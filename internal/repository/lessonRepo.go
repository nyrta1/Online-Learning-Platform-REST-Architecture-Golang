package repository

import (
	"gorm.io/gorm"
	"online-learning-platform/internal/models"
	"online-learning-platform/internal/rest/forms"
)

type LessonRepo interface {
	GetLessonByID(id uint) (*models.Lesson, error)
	GetAllLessons() ([]models.Lesson, error)
	DeleteLesson(id uint) error
	CreateLesson(course *models.Lesson) error
	UpdateLesson(id uint, lessonForm forms.LessonForm) error
}

type LessonRepository struct {
	db *gorm.DB
}

func NewLessonRepository(db *gorm.DB) *LessonRepository {
	return &LessonRepository{db}
}

func (lr *LessonRepository) GetLessonByID(id uint) (*models.Lesson, error) {
	var lesson models.Lesson
	if err := lr.db.First(&lesson, id).Error; err != nil {
		return nil, err
	}

	return &lesson, nil
}

func (lr *LessonRepository) GetAllLessons() ([]models.Lesson, error) {
	var lessons []models.Lesson
	if err := lr.db.Find(&lessons).Error; err != nil {
		return nil, err
	}

	return lessons, nil
}

func (lr *LessonRepository) DeleteLesson(id uint) error {
	var lesson models.Lesson
	if err := lr.db.First(&lesson, id).Error; err != nil {
		return err
	}

	if err := lr.db.Delete(&lesson).Error; err != nil {
		return err
	}

	return nil
}

func (lr *LessonRepository) CreateLesson(lesson *models.Lesson) error {
	if err := lr.db.Create(lesson).Error; err != nil {
		return err
	}

	return nil
}

func (lr *LessonRepository) UpdateLesson(id uint, lessonForm forms.LessonForm) error {
	var lesson models.Lesson
	if err := lr.db.First(&lesson, id).Error; err != nil {
		return err
	}

	lesson.Name = lessonForm.Name
	lesson.Description = lessonForm.Description
	lesson.VideoURL = lessonForm.VideoUrl

	if err := lr.db.Save(&lesson).Error; err != nil {
		return err
	}

	return nil
}
