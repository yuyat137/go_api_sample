package main

import (
  "fmt"
  "os"
)

func handleError(err error) {
  if err != nil {
    fmt.Println("Error:", err)
    os.Exit(1)
  }
}
