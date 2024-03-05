package repository

import (
	"gorm.io/gorm"
	"online-learning-platform/internal/models"
	"online-learning-platform/internal/rest/forms"
)

type CourseRepo interface {
	GetCourseByID(id uint) (*models.Course, error)
	GetAllCourses() ([]models.Course, error)
	DeleteCourse(id uint) error
	UpdateCourseByID(id uint, courseForm forms.CourseForm) error
	CreateCourse(course *models.Course) error
}

type CourseRepository struct {
	db *gorm.DB
}

func NewCourseRepository(db *gorm.DB) *CourseRepository {
	return &CourseRepository{db}
}

func (cr *CourseRepository) GetCourseByID(id uint) (*models.Course, error) {
	var course models.Course
	if err := cr.db.First(&course, id).Error; err != nil {
		return nil, err
	}

	return &course, nil
}

func (cr *CourseRepository) GetAllCourses() ([]models.Course, error) {
	var courses []models.Course
	if err := cr.db.Find(&courses).Error; err != nil {
		return nil, err
	}

	return courses, nil
}

func (cr *CourseRepository) DeleteCourse(id uint) error {
	var course models.Course
	if err := cr.db.First(&course, id).Error; err != nil {
		return err
	}

	if err := cr.db.Delete(&course).Error; err != nil {
		return err
	}

	return nil
}

func (cr *CourseRepository) CreateCourse(course *models.Course) error {
	if err := cr.db.Create(course).Error; err != nil {
		return err
	}

	return nil
}

func (cr *CourseRepository) UpdateCourseByID(id uint, courseForm forms.CourseForm) error {
	var course models.Course
	if err := cr.db.First(&course, id).Error; err != nil {
		return err
	}

	course.Name = courseForm.Name
	course.Description = courseForm.Description

	if err := cr.db.Save(&course).Error; err != nil {
		return err
	}

	return nil
}
