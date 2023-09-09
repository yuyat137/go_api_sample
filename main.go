package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql" //本コードでmysqlを使う記述が無いが、インポートしたいものに対しては"_"を頭につける決まり。
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

func setupRouter() *gin.Engine {
	router := gin.Default()

	dbInit()

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	router.GET("/users", func(c *gin.Context) {
			users := getAllUser()
			c.JSON(200, gin.H{"users": users})
	})

	return router
}

func main() {
	router := setupRouter()
	router.Run(":8080")
}
