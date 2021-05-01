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

  /* Create user folder to store logs, messages and
   * files */
   err := client.CreateDir()
   if err != nil { log.Fatal(err) }

  /* The client starts listening to its channel
   * waiting for messages to send */
  go client.SendMessages()

  /*  */
  go client.StartListening()

  for ; opt!= 'x'; {
    mainMenu()
    fmt.Scanf("%c\n", &opt)

    switch opt {
    case 'a':
      users, err := client.GetUsers()
      if err != nil {
        log.Println(err)
        continue
      }
      log.Println(users)

      dst := chooseUser(users, username)
      if dst == -1 { continue }

      client.SetDest(dst)
      client.SendChan <- getData(users[dst])
    case 'x':
      break
    default:
      log.Println("Invalid option")
    }

  }

  log.Println("Client disconnected")
}

func getData(dst string) string {
  reader := bufio.NewReader(os.Stdin)

  fmt.Printf(
    "Write the message to %s: (Press enter to send)\n",
    dst)
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

func chooseUser(users map[int]string, self string) int {
  var user string
  inverseMap := make(map[string]int)

  if len(users) <= 1 {
    log.Println("No users available")
    return -1
  }

  for k,v := range users {
    if v == self { continue }
    fmt.Println(" -", v)
    inverseMap[v] = k
  }

  fmt.Print("Write the name of the user to send message: ")
  fmt.Scanf("%s", &user)

  /* Try to get the user's pid from the map */
  ret, ok := inverseMap[user]
  if !ok {
    log.Println("Invalid username")
    return -1
  }

  return ret
}
