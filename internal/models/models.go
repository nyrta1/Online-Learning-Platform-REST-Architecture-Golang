package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID       int    `gorm:"primaryKey;autoIncrement" json:"ID,omitempty"`
	Name     string `gorm:"size:35" json:"name"`
	Surname  string `gorm:"size:35" json:"surname"`
	Username string `gorm:"type:varchar(35);unique_index" json:"username"`
	Password string `gorm:"type:varchar(255)" json:"password"`
	Roles    []Role `gorm:"many2many:user_roles"`
}

type Role struct {
	gorm.Model
	ID    uint   `gorm:"primaryKey;autoIncrement"`
	Name  string `gorm:"size:35"`
	Users []User `gorm:"many2many:user_roles"`
}

type Course struct {
	gorm.Model
	ID          uint     `gorm:"primaryKey;autoIncrement"`
	Name        string   `gorm:"size:35"`
	Description string   `gorm:"type:text"`
	Estimation  int      `gorm:"type:integer; default:0; max:5"`
	LessonCount int      `gorm:"type:integer"`
	Lessons     []Lesson `gorm:"many2many:course_lessons"`
	OwnerUserID int      `gorm:"type:integer"`
}

type Lesson struct {
	gorm.Model
	ID          uint     `gorm:"primaryKey;autoIncrement"`
	Name        string   `gorm:"size:35"`
	Description string   `gorm:"type:text"`
	VideoURL    string   `gorm:"type:text"`
	Course      []Course `gorm:"many2many:course_lessons"`
}
