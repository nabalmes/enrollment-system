package models

import (
	"time"

	"gorm.io/gorm"
)

type TimeTable struct {
	gorm.Model
	Subject   Subjects
	SubjectID uint
	StartTime time.Time
	EndTime   time.Time
	Day       time.Time
}
