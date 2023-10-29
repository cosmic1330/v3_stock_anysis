package models

import "time"

type DailyDeal struct {
	TransactionDate time.Time `gorm:"not null"`
	StockID     string `gorm:"type:varchar(10);not null"`
	StockName   string `gorm:"type:varchar(30);not null"`
	Volume      int `gorm:"not null"`
	OpenPrice   float64 `gorm:"type:numeric(10,2);not null"`
	ClosePrice  float64 `gorm:"type:numeric(10,2);not null"`
	HighPrice   float64 `gorm:"type:numeric(10,2);not null"`
	LowPrice    float64 `gorm:"type:numeric(10,2);not null"`
}
