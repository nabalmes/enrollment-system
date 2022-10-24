package models

import "gorm.io/gorm"

type Course struct {
	gorm.Model
	CourseCode  int
	Name        string
	Description string
}
