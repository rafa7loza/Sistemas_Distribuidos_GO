package labs

import (
  "os"
  "net"
  "log"
  "encoding/gob"
)

/* External functions */
func RegisterUser(username string) {
  /* Init connection */
  conn, err := net.Dial(PROTOCOL, PortString)
  if err != nil { log.Fatal(err) }

  /* Encode and send message */
  msg := NewMessage(REGISTER_CODE, os.Getpid(), username)
  err = gob.NewEncoder(conn).Encode(msg)

  /* Close connection */
  if err != nil { log.Println(err) }
  conn.Close()
}
