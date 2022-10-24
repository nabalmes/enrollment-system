package models

import "gorm.io/gorm"

type Users struct {
	gorm.Model
	Admin      string
	Cashier    string
	Registrar  string
	instructor string
}
