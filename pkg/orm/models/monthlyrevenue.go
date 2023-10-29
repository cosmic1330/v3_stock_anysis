package models

type MonthlyRevenue struct {
	Year                          int    `gorm:"not null"`
	Month                         int    `gorm:"not null"`
	StockID                       string `gorm:"type:varchar(10);not null"`
	CurrentMonthRevenue           int64
	PreviousMonthRevenue          int64
	PreviousYearSameMonthRevenue  int64
	MonthOverMonthRevenue         float64 `gorm:"type:numeric(15,3)"`
	YearOverYearRevenue           float64 `gorm:"type:numeric(15,3)"`
	CurrentYearCumulativeRevenue  int64
	PreviousYearCumulativeRevenue int64
	CompareCumulativeRevenue      float64 `gorm:"type:numeric(15,3)"`
}
