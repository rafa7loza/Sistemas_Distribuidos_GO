package labs

import (
  "os"
  "net"
  "log"
  "encoding/gob"
  "encoding/json"
  "time"
  "errors"
)

type Client struct {
  username string
  SendChan chan string
}

type client interface {
  GetUser() string
  RegisterUser()
  GetUsers() (map[int]string, error)
}


/* External functions */
func NewClient(name string) * Client {
  /* Create and init channel */
  ch := make(chan string, 1)
  ch <- ""
  return  &Client{name, ch}

}

/* Client methods */
func (c * Client) RegisterUser() {
  /* Init connection */
  conn, err := net.Dial(PROTOCOL, ADDRESS)
  if err != nil { log.Fatal(err) }

  /* Encode and send message */
  msg := NewMessage(REGISTER_CODE, os.Getpid(), c.username)
  err = gob.NewEncoder(conn).Encode(msg)
  if err != nil { log.Println(err) }

  /* Close connection */
  conn.Close()
}

func (c * Client) SendMessages() {
  for {
    select {
    case <-c.SendChan:
      for {
        conn, err := net.Dial(PROTOCOL, ADDRESS)
        if err != nil { log.Fatal(err) }

        /* Get data from the channel */
        data, ok := <- c.SendChan
        if !ok { break }

        /* Write message to the server */
        msg := NewMessage(0, os.Getpid(), data)
        err = gob.NewEncoder(conn).Encode(msg)

        if err != nil { log.Println(err) }
        conn.Close()
        time.Sleep(time.Millisecond * WAIT_TIME_MS)
      }

    default:
      time.Sleep(time.Millisecond * WAIT_TIME_MS)
    }
  }
}

func (c * Client) GetUsers() (map[int]string, error) {
  /* Init connection */
  conn, err := net.Dial(PROTOCOL, ADDRESS)
  if err != nil { log.Fatal(err) }
  defer conn.Close()

  /* Encode and send message */
  msg := NewMessage(GETUSERS_CODE, os.Getpid(), "")
  err = gob.NewEncoder(conn).Encode(msg)
  if err != nil { return nil, err }

  /* Listen for reply */
  err = gob.NewDecoder(conn).Decode(&msg)
  if err != nil { return nil, err }
  if msg.Code != GETUSERS_CODE { return nil, errors.New("Invalid code") }

  /* Return a map of users */
  var users map[int]string
  err = json.Unmarshal(msg.Data, &users)
  if err != nil { return nil, err }

  return users, nil
}
