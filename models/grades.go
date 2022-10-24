package models

import "gorm.io/gorm"

type Grades struct {
	gorm.Model
	Students   Students
	StudentsID uint
	Course     Course
	CourseID   uint
	Users      Users
	UsersID    uint
}
