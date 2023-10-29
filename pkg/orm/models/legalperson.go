package models

import "time"

type LegalPerson struct {
	TransactionDate  time.Time `gorm:"type:date;primaryKey"`
	StockID          string    `gorm:"type:varchar(10);not null;primaryKey"`
	StockName        string    `gorm:"type:varchar(30);not null"`
	ForeignInvestors int       `gorm:"not null"`
	InvestmentTrust  int       `gorm:"not null"`
	Dealer           int       `gorm:"not null"`
}
