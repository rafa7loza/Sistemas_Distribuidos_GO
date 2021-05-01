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

/*
func initServer() {
  users = make(map[int]string)
  ln, err := net.Listen("tcp", labs.PortString)
  defer ln.Close()

  if err != nil {
    log.Println(err, "initServer")
    return
  }

  for {

    conn, err := ln.Accept()
    if err != nil { log.Println(err, "ln.Accept()") }
    go labs.HandleClient(conn, users)
    fmt.Println(users)
  }
}
*/
