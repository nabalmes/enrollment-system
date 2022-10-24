package models

import (
	"time"

	"gorm.io/gorm"
)

type TransactionLine struct {
	gorm.Model
	Transaction			Transactions
	TransactionID		int
	PaymentType			string
	Period				time.Time
	TotalAmount			int
	Balance				int
	DatePaid			time.Time
	PaidTime			time.Time
	VerifiedBy			string
	DateCreated			time.Time
	UpdatedBy			string
}
