package labs

import (
  "os"
  "net"
  "log"
  "encoding/gob"
)

type Client struct {
  username string
}

type client interface {
  GetUser() string
  RegisterUser()
}


/* External functions */
func NewClient(name string) * Client {
  return &Client{name}
}

/* Client methods */
func (c * Client) RegisterUser() {
  /* Init connection */
  conn, err := net.Dial(PROTOCOL, PortString)
  if err != nil { log.Fatal(err) }

  /* Encode and send message */
  msg := NewMessage(REGISTER_CODE, os.Getpid(), c.username)
  err = gob.NewEncoder(conn).Encode(msg)

  /* Close connection */
  if err != nil { log.Println(err) }
  conn.Close()
}
