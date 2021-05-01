package main

import(
  "fmt"
  "net"
  "os"
  "bufio"
  "labs"
  "time"
  "log"
  "encoding/gob"
  "strings"
)

func main() {
  var opt rune
  var username string

  /* Create and init channel */
  ch := make(chan string, 1)
  ch <- ""
  go sendData(ch)

  /* Register username in server */
  fmt.Print("Enter your username: ")
  fmt.Scanf("%s", &username)
  client := labs.NewClient(username)
  client.RegisterUser()

  for ; opt!= 'x'; {
    mainMenu()
    fmt.Scanf("%c\n", &opt)

    switch opt {
    case 'a':
      ch <- getData()
    case 'x':
      break
    default:
      log.Println("Invalid option")
    }

  }

  log.Println("Client disconnected")
}

func sendData(inputData chan string) {

  for {
    select {
    case <-inputData:
      for {
        conn, err := net.Dial("tcp", labs.PortString)
        if err != nil { log.Fatal(err) }

        /* Get data from the channel */
        data, ok := <- inputData
        if !ok { break }

        /* Write message to the server */
        msg := labs.NewMessage(0, os.Getpid(), data)
        err = gob.NewEncoder(conn).Encode(msg)

        if err != nil { log.Println(err) }
        conn.Close()
        time.Sleep(time.Millisecond * labs.WAIT_TIME_MS)
      }

    default:
      time.Sleep(time.Millisecond * labs.WAIT_TIME_MS)
    }
  }
}

func getData() string {
  reader := bufio.NewReader(os.Stdin)

  data, err := reader.ReadString('\n')
  if err != nil {
    return "(NULL)"
  }
  data = strings.Replace(data, "\n", "", -1)
  log.Println("data", data)
  return data
}

func mainMenu() {
  fmt.Println("a) Send message")
  fmt.Println("b) Send file")
  fmt.Println("c) Show messages received")
  fmt.Println("x) Exit\n")
  fmt.Print("Enter an option: ")
}
