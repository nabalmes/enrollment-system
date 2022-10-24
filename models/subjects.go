package models

import "gorm.io/gorm"

type Subjects struct {
	gorm.Model
	Code         int
	Name         string
	Course       string
	User         Users
	UsersID      uint
	Description  string
	SchoolYear   SchoolYear
	SchoolYearID uint
}
