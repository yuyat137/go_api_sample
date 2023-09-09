package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql" //本コードでmysqlを使う記述が無いが、インポートしたいものに対しては"_"を頭につける決まり。
)

func setupRouter() *gin.Engine {
	router := gin.Default()

	dbInit()

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
