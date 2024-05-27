package mysql

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Init() {
	host := "127.0.0.1"
	post := "3306"
	user := "root"
	pass := "123"
	db_name := "emsvc"
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc",
		user, pass, host, post, db_name)
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN: dsn,
	}))
	if err != nil {
		panic("failed to connect database")
	}
	DB = db
}
