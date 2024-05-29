package NFT

import (
	"Microservice/server/ent"
	"context"
	"fmt"
	"log"

	"entgo.io/ent/dialect"
	_ "github.com/go-sql-driver/mysql"
)

var Mysql *ent.Client

func initTable() {
	ctx := context.Background()
	// Run the automatic migration tool to create all schema resources.
	if err := Mysql.Schema.Create(ctx); err != nil {
		log.Printf("failed creating schema resources: %v", err)
	}
	log.Printf("created Table Trasaction")
}

func InitMysql() {
	host := "localhost"
	port := "3306"
	user := "root"
	pass := "123"
	db_name := "emsvc"
	//dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc",
	//user, pass, host, post, db_name)
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		user, pass, host, port, db_name)

	db, err := ent.Open(dialect.MySQL, dsn)
	if err != nil {
		panic("failed to connect database err : " + err.Error())
	}
	Mysql = db
	initTable()
}
