package main

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
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

func getUser(id int) User {
  db, err := gormConnect()

  if err != nil {
    fmt.Println("Error:", err)
    os.Exit(1)
  }

  var user User
  db.First(&user, id)
  defer db.Close()
  return user
}
