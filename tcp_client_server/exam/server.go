package main

import (
  "fmt"
  "log"
  "net"
  "labs"
  "encoding/gob"
)

var pids map[int]bool

func main() {
  log.Println("Starting server...")
  go initServer()

  var input string
  fmt.Scanln(&input)
  log.Println("Server stopped")
}

func initServer() {
  ln, err := net.Listen("tcp", labs.PortString)
  defer ln.Close()

  if err != nil {
    log.Println(err, "initServer")
    return
  }

  for {
    /* accept connection on port */
    conn, err := ln.Accept()
    if err != nil { log.Println(err, "ln.Accept()") }
    go handleClient(conn)
  }

}

func handleClient(conn net.Conn) {
  var msg labs.Message

  err := gob.NewDecoder(conn).Decode(&msg)

  if err != nil {
    log.Println(err)
    return
  }

  fmt.Println("Message=", msg)
  conn.Close()
}
