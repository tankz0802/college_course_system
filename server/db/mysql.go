package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql" //引入数据库连接包
)

//DB 数据库连接句柄
var (
	DB  *sql.DB
	err error
)

//Init 初始化数据库连接
func Init() {

	sqlInfo := "root:123456@tcp(106.52.211.246:3306)/database3?charset=utf8"
	DB, err = sql.Open("mysql", sqlInfo)
	if err != nil {
		log.Fatal(err.Error())
	}

	DB.SetConnMaxLifetime(100)
	DB.SetMaxIdleConns(10)
	if err := DB.Ping(); err != nil {
		fmt.Println("Mysql Failed Connected!")
		log.Fatal(err.Error())
	}
	fmt.Println("Mysql Successfully Connected!")
}
