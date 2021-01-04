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
	cfg, err := ini.Load("config.ini")
	if err != nil {
		fmt.Printf("Fail to read file: %v", err)
		os.Exit(1)
	}

	mysqlSection := cfg.Section("mysql")
	MYSQL_HOST = mysqlSection.Key("host").String()
	MYSQL_PORT, _ = mysqlSection.Key("port").Int64()
	MYSQL_USER = mysqlSection.Key("user").String()
	MYSQL_PASSWORD = mysqlSection.Key("password").String()
	DATABASE = mysqlSection.Key("database").String()
}

