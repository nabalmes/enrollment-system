package models

import (
	"time"

	"github.com/jinzhu/gorm"
	_ "gorm.io/gorm"
)

type Students struct {
	gorm.Model

	FirstName     string
	MiddleName    string
	LastName      string
	Age           int
	Gender        string
	Address       string
	Birthday      time.Time
	BirthdayPlace string
	Status        string
	Nationality   string
	Religion      string
	RegilionID    uint
	MobileNumber  int
	Email         string
}
