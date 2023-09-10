package main

import (
  "github.com/jinzhu/gorm"
)

func gormConnect() (*gorm.DB, error) {
  DBMS := "mysql"
  USER := "test"
  PASS := "12345678"
  DBNAME := "go_api_sample"
  // MySQLだと文字コードの問題で"?parseTime=true"を末尾につける必要がある
  CONNECT := USER + ":" + PASS + "@/" + DBNAME + "?parseTime=true"
  db, err := gorm.Open(DBMS, CONNECT)

  return db, err
}

func dbInit() {
  db, err := gormConnect()
  handleError(err)
  // コネクション解放
  defer db.Close() // defer: 関数が終わった時に必ず実行する
  db.AutoMigrate(&User{}) //構造体に基づいてテーブルを作成
}

func dbInsert(record interface{}) {
  db, err := gormConnect()
  handleError(err)

  defer db.Close()
  db.Create(record)
}

