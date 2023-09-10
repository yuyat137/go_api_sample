package main

import (
  "github.com/gin-gonic/gin"
  "strconv"
)

func getAllUserHandler(c *gin.Context) {
  c.JSON(200, gin.H{"users": getAllUser()})
}

func getUserHandler(c *gin.Context) {
  id_string := c.Param("id")
  id, _ := strconv.Atoi(id_string)
  c.JSON(200, gin.H{ "user": getUser(id) })
}

func postUserHandler(c *gin.Context) {
  name := c.PostForm("name")
  password := c.PostForm("password")
  createUser(name, password)
  c.JSON(204, nil)
}
