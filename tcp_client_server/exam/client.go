package main

import(
  "fmt"
  "os"
  "bufio"
  "labs"
  "log"
  "strings"
)

func main() {
  var opt rune
  var username string

  /* Register username in server */
  fmt.Print("Enter your username: ")
  fmt.Scanf("%s", &username)
  client := labs.NewClient(username)
  client.RegisterUser()

  /* The client starts listening to its channel
   * waiting for messages to send */
  go client.SendMessages()

  for ; opt!= 'x'; {
    mainMenu()
    fmt.Scanf("%c\n", &opt)

    switch opt {
    case 'a':
      client.SendChan <- getData()
    case 'x':
      break
    default:
      log.Println("Invalid option")
    }

  }

  log.Println("Client disconnected")
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
