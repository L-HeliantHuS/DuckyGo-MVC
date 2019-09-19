package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"os"
	"time"
)

var DB *gorm.DB

func Init() {

	db, err := gorm.Open("mysql", os.Getenv("MYSQL_DSN"))

	// 开启Log
	db.LogMode(true)
	if err != nil {
		panic(err)
	}

	//设置连接池
	//空闲
	db.DB().SetMaxIdleConns(300)
	//打开
	db.DB().SetMaxOpenConns(500)
	//超时
	db.DB().SetConnMaxLifetime(time.Second * 30)

	DB = db

}
