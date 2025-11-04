package repository

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"medicine/config"
)

var (
	MysqlClient *sql.DB
)

func InitMysql() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s",
		config.GlobalConfig.Mysql.Username,
		config.GlobalConfig.Mysql.Password,
		config.GlobalConfig.Mysql.Addr,
		config.GlobalConfig.Mysql.Port,
		config.GlobalConfig.Mysql.Databases,
		config.GlobalConfig.Mysql.Charset)
	var err error
	MysqlClient, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("打开 MySQL 连接失败: %v", err)
	}
	if err = MysqlClient.Ping(); err != nil {
		log.Fatalf("连接 MySQL 失败: %v", err)
	}
	log.Println("MySQL 连接成功")
}
