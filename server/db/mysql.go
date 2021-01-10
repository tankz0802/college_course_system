package db

import (
	"ccs/config"
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

	sqlInfo := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?timeout=15s&readTimeout=15s&writeTimeout=15s&charset=utf8",
		config.MYSQL_USER,
		config.MYSQL_PASSWORD,
		config.MYSQL_HOST,
		config.MYSQL_PORT,
		config.DATABASE)
	DB, err = sql.Open("mysql", sqlInfo)
	if err != nil {
		log.Fatal(err.Error())
	}

	DB.SetConnMaxLifetime(100)
	//DB.SetMaxIdleConns(10)
	DB.SetMaxIdleConns(0)
	if err := DB.Ping(); err != nil {
		fmt.Println("Mysql Failed Connected!")
		log.Fatal(err.Error())
	}
	fmt.Println("Mysql Successfully Connected!")
}
