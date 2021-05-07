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

    // TODO: Refactor this code into a function in order to
    // avoid code repetition
    err = gob.NewDecoder(conn).Decode(&msg)
    if err != nil {
      if err != io.EOF { log.Println(err) }
    }

    switch msg.Code {
    case RECMESSAGE_CODE:
      c.saveMessage(msg.Data)
    case RECFILE_CODE:
      c.saveFile(msg.Data)
    }

    conn.Close()
    time.Sleep(time.Millisecond * WAIT_TIME_MS)
  }
}

func (c * Client) SendFile(name string) error {
  path := c.Folder + "/" + name

  /* Open the file read-only mode */
  fd, err := os.Open(path)
  if err != nil { return err }
  defer fd.Close()

  /* Get the size of the file  and rezise the byte slice */
  st, err := fd.Stat()
  if err != nil { return err }
  data := make([]byte, st.Size())

  /* Fill the slice with the file's content */
  _, err = fd.Read(data)
  if err != nil { return err }

  /* Connect to the server */
  conn, err := net.Dial(PROTOCOL, ADDRESS)
  if err != nil { return err }
  defer conn.Close()

  /* Write message to the server */
  msg := NewMessage(
    SENDFILE_CODE,
    os.Getpid(),
    c.Dest,
    File{name, data})
  err = gob.NewEncoder(conn).Encode(msg)
  if err != nil { return err }

  /* No error */
  return nil
}

func (c * Client) GetFiles() ([]string, error) {
  ret := make([]string, 0)

  files, err := os.ReadDir(c.Folder)
  if err != nil { return nil, err }

  for _, file := range files {
    if c.username + ".msg" == file.Name() { continue }
    ret = append(ret, file.Name())
  }

  return ret, nil
}

func (c * Client) Disconnect() {
  /* Connect to the server */
  conn, err := net.Dial(PROTOCOL, ADDRESS)
  if err != nil { log.Println(err) }
  defer conn.Close()

  /* Write message to the server */
  msg := NewMessage(
    DISCONNECT_CODE,
    os.Getpid(),
    -1,
    []byte{})
  err = gob.NewEncoder(conn).Encode(msg)
  if err != nil { log.Println(err) }

}

func (c * Client) GetLogs() string {
  filename := c.Folder + "/" + c.username + ".msg"
  data, err := os.ReadFile(filename)
  if err != nil {
    log.Println(err)
    return ""
  }

  return string(data)
}
/* Private methods */
func (c * Client) saveMessage(data []byte) {
  /* Append the message to the file */
  filename := c.Folder + "/" + c.username + ".msg"
  if err := AppendToFile(filename, data); err != nil {
    log.Println(err)
  }
}

func (c * Client) saveFile(data []byte) error {
  var file File

  /* Decode the File struct */
  err := json.Unmarshal(data, &file)
  if err != nil { return err }

  filename := c.Folder + "/" + file.Name
  fd, err := os.OpenFile(
    filename,
    os.O_CREATE|os.O_WRONLY,
    0644)

  if err != nil { return err }
  defer fd.Close()

  _, err = fd.Write(file.Content)
  if err != nil { return err }

  return nil
}
