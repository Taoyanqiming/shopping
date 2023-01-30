package dao

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func SqlInit() (err error) {
	db, err = sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/shopping?charset=utf8")
	if err != nil {
		return
	}
	err = db.Ping()
	if err != nil {
		return
	}
	db.SetMaxOpenConns(100)
	db.SetMaxIdleConns(10)
	return
}

func SqlConn() *sql.DB {
	err := SqlInit()
	if err != nil {
		fmt.Printf("数据库连接失败！, err:%v\n", err)
		return db
	}
	fmt.Println("数据库连接成功！")
	return db
}
