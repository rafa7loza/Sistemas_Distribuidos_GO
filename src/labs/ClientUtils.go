package labs

import (
  "io"
  "os"
  "net"
  "log"
  "fmt"
  "encoding/gob"
  "encoding/json"
  "time"
  "errors"
)

type Client struct {
  username string
  SendChan chan string
  Dest int
  Folder string
}

type client interface {
  GetUser() string
  RegisterUser()
  GetUsers() (map[int]string, error)
  SetDest(dst int)
  StartListening() error
}

/* Internal functions */
func badCode() error {
  return errors.New("Invalid code")
}

func appendToFile(filename string, msg []byte) error {
  fd, err := os.OpenFile(
    filename,
    os.O_APPEND|os.O_CREATE|os.O_WRONLY,
    0644)

  if err != nil { return nil }

  defer fd.Close()
  if _, err := fd.Write(msg); err != nil { return err }
  return nil
}

/* External functions */
func NewClient(name string) * Client {
  /* Create and init channel */
  ch := make(chan string, 1)
  ch <- ""
  return  &Client{name, ch, -1, ""}

}

/* Client methods */
func (c * Client) RegisterUser() {
  /* Init connection */
  conn, err := net.Dial(PROTOCOL, ADDRESS)
  if err != nil { log.Fatal(err) }

  /* Encode and send message */
  msg := NewMessage(REGISTER_CODE, os.Getpid(), -1, c.username)
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
        msg := NewMessage(
          SENDMESSAGE_CODE,
          os.Getpid(),
          c.Dest,
          data)
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
  msg := NewMessage(GETUSERS_CODE, os.Getpid(), -1, "")
  err = gob.NewEncoder(conn).Encode(msg)
  if err != nil { return nil, err }

  /* Listen for reply */
  err = gob.NewDecoder(conn).Decode(&msg)
  if err != nil { return nil, err }
  if msg.Code != GETUSERS_CODE { return nil, badCode() }

  /* Return a map of users */
  var users map[int]string
  err = json.Unmarshal(msg.Data, &users)
  if err != nil { return nil, err }

  return users, nil
}

func (c * Client) SetDest(dst int) {
  c.Dest = dst
}

func (c * Client) CreateDir() error {
  path, err := os.Getwd()
  if err != nil { return err }

  folder := fmt.Sprintf("%s_%d", c.username, os.Getpid())
  path += "/clients/" + folder

  err = os.MkdirAll(path, 0777)
  if err != nil { return err }

  /* Set the path atributte */
  c.Folder = path
  return nil
}

func (c * Client) StartListening()  {
  for {
    /* Init connection */
    conn, err := net.Dial(PROTOCOL, ADDRESS)
    if err != nil { log.Fatal(err) }

    /* Encode and send message */
    msg := NewMessage(CHECKMSG_CODE, os.Getpid(), -1, "")
    err = gob.NewEncoder(conn).Encode(msg)
    if err != nil {
      log.Println(err)
      continue
    }

    c.saveMessage(conn)

    conn.Close()
    time.Sleep(time.Millisecond * WAIT_TIME_MS)
  }
}

/* Private methods */
func (c * Client) saveMessage(conn net.Conn) {
  var msg Message

  // TODO: Refactor this code into a function in order to
  // avoid code repetition
  err := gob.NewDecoder(conn).Decode(&msg)
  if err != nil {
    if err != io.EOF { log.Println(err) }
    return
  }

  if msg.Code != RECMESSAGE_CODE {
    log.Println(badCode())
    return
  }

  /* Append the message to the file */
  filename := c.Folder + "/" + c.username + ".msg"
  if err := appendToFile(filename, msg.Data); err != nil {
    log.Println(err)
  }
}
