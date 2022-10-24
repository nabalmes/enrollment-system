package models

import "gorm.io/gorm"

type Transactions struct {
	gorm.Model
	ReferenceNumber int
	Students        Students
	StudentsID      uint
	Course          Course
	CourseID        uint
	SchoolYear      SchoolYear
	SchoolYearID    uint
	Balance         int
}
