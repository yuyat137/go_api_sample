package main

import (
  "github.com/gin-gonic/gin"
)

func getAllUserHandler(c *gin.Context) {
  c.JSON(200, gin.H{"users": getAllUser()})
}
