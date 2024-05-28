package NFT

import (
	"Microservice/ent"
	"fmt"

	"entgo.io/ent/dialect"
)

var Mysql *ent.Client

func InitMysql() {
	host := "127.0.0.1"
	post := "3306"
	user := "root"
	pass := "123"
	db_name := "emsvc"
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc",
		user, pass, host, post, db_name)
	db, err := ent.Open(dialect.MySQL, dsn)
	if err != nil {
		panic("failed to connect database")
	}
	Mysql = db
}
