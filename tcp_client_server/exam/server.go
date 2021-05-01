package main

import (
  "fmt"
  "log"
  "labs"
)

var users map[int]string

func main() {
  log.Println("Starting server...")
  var srvr labs.Server

  go srvr.Init()

  var input string
  fmt.Scanln(&input)
  log.Println("Server stopped")
}
