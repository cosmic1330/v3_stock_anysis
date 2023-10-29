package models

type Eps struct {
	Season             string  `gorm:"not null"`
	StockID            string  `gorm:"type:varchar(10);not null"`
	StockName          string  `gorm:"type:varchar(30);not null"`
	EpsData            float64 `gorm:"type:numeric(10,3);not null"`
	Revenue            int64   `gorm:"not null"`
	OperatingIncome    int64   `gorm:"not null"`
	NonOperatingIncome int64   `gorm:"not null"`
	PreTaxIncome       int64   `gorm:"not null"`
	NetIncome          int64   `gorm:"not null"`
}
