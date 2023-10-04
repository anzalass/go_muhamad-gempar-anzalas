package model

import (
	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitModel() *gorm.DB {
	dsn := "root:@tcp(localhost:3306)/db_unittest?parseTime=true"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		logrus.Error("Cant Connect to database")
	}

	return db
}

func Migrate(db *gorm.DB) {
	db.AutoMigrate(&User{})
}
