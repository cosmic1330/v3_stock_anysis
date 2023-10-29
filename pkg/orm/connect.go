package orm

import (
	"fmt"
	"stockPredict/pkg/orm/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Database *gorm.DB
var err error

func Connect() {
	dsn := "host=localhost user=root password=secret dbname=postgres port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	Database, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	// 自动创建表格
	err = Database.AutoMigrate(
		&models.MonthlyRevenue{},
		&models.DailyDeal{},
		&models.Date{},
		&models.EpsSeason{},
		&models.Eps{},
		&models.Stock{},
		&models.LegalPerson{})
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("success connect database")
}
