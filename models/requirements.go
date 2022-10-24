package models

import (
	"gorm.io/gorm"
)

type Requirements struct {
	gorm.Model
	Students   Students
	StudentsID uint
	PSA        bool
	Baptismal  bool
	TOR        bool
}
