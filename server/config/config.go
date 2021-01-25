package config

import (
	"fmt"
	"gopkg.in/ini.v1"
	"os"
)

var (
	MYSQL_HOST string
	MYSQL_PORT int64
	MYSQL_USER string
	MYSQL_PASSWORD string
	DATABASE string
)

func Init() {
	cfg, err := ini.Load("config/config.ini")
	if err != nil {
		fmt.Printf("Fail to read file: %v", err)
		os.Exit(1)
	}

	mysqlSection := cfg.Section("mysql")
	MYSQL_HOST = mysqlSection.Key("host").MustString("127.0.0.1")
	MYSQL_PORT = mysqlSection.Key("port").MustInt64(3306)
	MYSQL_USER = mysqlSection.Key("user").MustString("root")
	MYSQL_PASSWORD = mysqlSection.Key("password").MustString("123456")
	DATABASE = mysqlSection.Key("database").MustString("ccs")
}

