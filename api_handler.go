package main

import (
  "github.com/gin-gonic/gin"
)

func getAllUserHandler(c *gin.Context) {
  c.JSON(200, gin.H{"users": getAllUser()})
}

func postUserHandler(c *gin.Context) {
  name := c.PostForm("name")
  password := c.PostForm("password")
  createUser(name, password)
  c.JSON(204, nil)
}
