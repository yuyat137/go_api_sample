package main

import (
  "github.com/jinzhu/gorm"
  "fmt"
  "os"
)

// Userモデル宣言
// モデルはDBのテーブル構造をGOの構造体で表したもの
type User struct {
  gorm.Model // CreatedAt, UpdatedAt, DeletedAtを自動で入れてくれる
  Name string `form:"name" binding:"required"`
}

func getAllUser() []User {
  db, err := gormConnect()

  if err != nil {
    fmt.Println("Error:", err)
    os.Exit(1)
  }

  defer db.Close()
  var users []User
  db.Order("created_at desc").Find(&users)
  return users
}

func createUser(name string) {
  user := User{Name: name}
  dbInsert(&user)
}
