package models

import "time"

type Date struct {
	TransactionDate time.Time `gorm:"type:date;primaryKey"`
}
