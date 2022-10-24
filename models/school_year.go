package models

import (
	"time"

	"gorm.io/gorm"
)

type SchoolYear struct {
	gorm.Model
	Year     time.Time
	Semester int
}
