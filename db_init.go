package main

import (
	"github.com/jinzhu/gorm"
)

func gormConnect() *gorm.DB {
	DBMS := "mysql"
	USER := "test"
	PASS := "12345678"
	DBNAME := "go_api_sample"
	// MySQLだと文字コードの問題で"?parseTime=true"を末尾につける必要がある
	CONNECT := USER + ":" + PASS + "@/" + DBNAME + "?parseTime=true"
	db, err := gorm.Open(DBMS, CONNECT)

	if err != nil {
		panic(err.Error())
	}
	return db
}

func dbInit() {
	db := gormConnect()

	// コネクション解放
	defer db.Close() // defer: 関数が終わった時に必ず実行する
	db.AutoMigrate(&User{}) //構造体に基づいてテーブルを作成
}

