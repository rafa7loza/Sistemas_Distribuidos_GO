package main

import (
  "fmt"
  "log"
  "net"
  "net/rpc"
  "labs"
  "labs/rpclab"
)

func server() {
  rpc.Register(new(rpclab.Server))

  ln, err := net.Listen("tcp", ":"+labs.PORT)
  if err != nil { log.Fatal(err) }
  log.Println("Iniciando el servidor")

  for {
    conn, err := ln.Accept()
    if err != nil {
      log.Println("Error on ln.Accept():", err)
      continue
    }
    go rpc.ServeConn(conn)
  }
}

func main() {
  go server()

  var input string
  fmt.Scanln(&input)
}
