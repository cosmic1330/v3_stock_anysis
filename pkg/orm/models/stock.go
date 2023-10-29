package models

type Stock struct {
	StockID    string `gorm:"type:varchar(10);primaryKey"  json:"stock_id"`
	StockName  string `gorm:"type:varchar(30);not null"  json:"stock_name"`
	Enabled    bool   `gorm:"default:true"  json:"enabled"`
}
// TableName sets the insert table name for this struct type
func (Stock) TableName() string {
    return "stock"
}