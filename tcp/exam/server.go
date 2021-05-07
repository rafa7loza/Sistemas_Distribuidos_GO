package main

import (
  "fmt"
  "log"
  "labs"
)

func main() {
  var srvr labs.Server
  var input rune

  /* Start the server */
  log.Println("Starting server...")
  go srvr.Init()

  for ; input != 'x'; {
    menu()
    fmt.Scanf("%c\n", &input)

    switch input {
    case 'a':
      srvr.PrintLogs()
    case 'b':
      srvr.FlushLogs()
      log.Println("Logs stored in file")
    case 'x':
      log.Println("Stopping the service")
    default:
      log.Println("Invalid option")
    }
  }
  log.Println("Server stopped")
}

func menu() {
  fmt.Println("a) Show logs")
  fmt.Println("b) Backup logs")
  fmt.Println("x) Stop service")
  fmt.Print("Choose an option: ")
}
