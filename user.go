package main

import (
	"github.com/jinzhu/gorm"
)

// Userモデル宣言
// モデルはDBのテーブル構造をGOの構造体で表したもの
type User struct {
	gorm.Model // CreatedAt, UpdatedAt, DeletedAtを自動で入れてくれる
	Name string `form:"name" binding:"required"`
}

func getAllUser() []User {
	db := gormConnect()

	defer db.Close()
	var users []User
	db.Order("created_at desc").Find(&users)
	return users
}
