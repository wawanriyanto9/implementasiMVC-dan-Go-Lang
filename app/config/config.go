package config

import (
	"github.com/imadedwisuryadinata/implementasiMvcGolang/app/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func DBInit() *gorm.DB {
	db, err := gorm.Open(mysql.Open("root:1234@/simple_bank?charset=utf8&parseTime=True&loc=Local"), &gorm.Config{})
	if err != nil {
		panic("failed to connect to database" + err.Error())
	}
	db.AutoMigrate(new(model.Account), new(model.Transaction))
	return db
}
