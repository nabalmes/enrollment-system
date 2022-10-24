package models

import "gorm.io/gorm"

type Fees struct {
	gorm.Model
	Course        Course
	CourseID      uint
	LibraryFee    int
	Miscellaneous float64
}
