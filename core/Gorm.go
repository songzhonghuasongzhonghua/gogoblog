package core

import (
	"fmt"
	"github.com/songzhonghuasongzhonghua/gogoblog/global"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"strconv"
)

func InitGorm() *gorm.DB {
	mysqlConf := global.Config.Mysql
	dsn := mysqlConf.User + ":" + mysqlConf.Password + "@tcp(" + mysqlConf.Host + ":" + strconv.Itoa(mysqlConf.Port) + ")/" + mysqlConf.DB

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("链接成功")
	return db
}
