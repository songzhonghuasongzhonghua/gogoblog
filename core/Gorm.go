package core

import (
	"github.com/songzhonghuasongzhonghua/gogoblog/global"
	"github.com/songzhonghuasongzhonghua/gogoblog/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"strconv"
)

func InitGorm() *gorm.DB {
	mysqlConf := global.Config.Mysql
	dsn := mysqlConf.User + ":" + mysqlConf.Password + "@tcp(" + mysqlConf.Host + ":" + strconv.Itoa(mysqlConf.Port) + ")/" + mysqlConf.DB + mysqlConf.Config

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	err = db.AutoMigrate(&models.BannerModel{})
	if err != nil {
		log.Fatal(err)
	}
	return db
}
