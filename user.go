package main

import (
  "github.com/jinzhu/gorm"
  "golang.org/x/crypto/bcrypt"
)

// Userモデル宣言
// モデルはDBのテーブル構造をGOの構造体で表したもの
type User struct {
  gorm.Model // CreatedAt, UpdatedAt, DeletedAtを自動で入れてくれる
  Name string `form:"name" binding:"required"`
  Password string `form:"password" binding:"required"`
}

func getAllUser() []User {
  db, err := gormConnect()

  handleError(err)

  defer db.Close()
  var users []User
  db.Order("created_at desc").Find(&users)
  return users
}

func createUser(name string, password string) {
  hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
  if err != nil {
      // handle error
      return
  }
  user := User{
    Name: name,
    Password: string(hashedPassword),
  }
  dbInsert(&user)
}
