package labs

import (
  "io"
  "log"
  "net"
  "fmt"
  "time"
  "encoding/gob"
)

type Server struct {
  ln net.Listener
  Users map[int]string
  PendingMsg map[int][]Message
}

type server interface {
  Init()
  Run()
  HandleClient(conn net.Conn)
}

/* Server methods */
func (s * Server) Init() {
  s.Users = make(map[int]string)
  s.PendingMsg = make(map[int][]Message)

  ln, err := net.Listen(PROTOCOL, ADDRESS)
  s.ln = ln
  defer s.ln.Close()

  if err != nil {
    log.Println(err, "initServer")
    return
  }

  s.Run()
}

func (s * Server) Run() {
  for {
    /* Accept connection on port */
    conn, err := s.ln.Accept()
    if err != nil { log.Println(err, "ln.Accept()") }
    go s.HandleClient(conn)
  }
}

func (s * Server) HandleClient(conn net.Conn) {
  var msg Message
  defer conn.Close()

  // TODO: Refactor this code into a function in order to
  // avoid code repetition
  err := gob.NewDecoder(conn).Decode(&msg)
  if err != nil {
    if err != io.EOF { log.Println(err) }
    return
  }

  switch msg.Code {
  case REGISTER_CODE: // Add user to the map
    s.Users[msg.Id] = string(msg.Data)
    log.Println("User registered")

  case GETUSERS_CODE: // Send the map of users
    msg = *NewMessage(GETUSERS_CODE, -1, -1, s.Users)
    err = gob.NewEncoder(conn).Encode(msg)
    if err != nil {
      if err != io.EOF { log.Println(err) }
      log.Println(err)
      return
    }

  case SENDMESSAGE_CODE: // Send message to the other client
    /* Format message */
    dt := time.Now()
    text := fmt.Sprintf(
      "%s: %s wrote '%s'\n",
      dt.Format("01-02-2006 15:04:05"),
      s.Users[msg.Id],
      string(msg.Data))

    /* Create a message with the text and add it to the waiting list */
    msg = *NewMessage(RECMESSAGE_CODE, msg.Id, msg.Dest, []byte(text))
    s.PendingMsg[msg.Dest] = append(s.PendingMsg[msg.Dest], msg)

    /* Return the message */
    msg = *NewMessage(OK_CODE, -1, -1, []byte{})
    gob.NewEncoder(conn).Encode(msg)

  case SENDFILE_CODE:
    /* Create a message with the text and add it to the waiting list */
    msg = *NewMessage(RECFILE_CODE, msg.Id, msg.Dest, msg.Data)
    s.PendingMsg[msg.Dest] = append(s.PendingMsg[msg.Dest], msg)

    /* Return the message */
    msg = *NewMessage(OK_CODE, -1, -1, []byte{})
    gob.NewEncoder(conn).Encode(msg)

  case CHECKMSG_CODE: // Check if the client has pending messages to serve
    if _, ok := s.PendingMsg[msg.Id]; !ok { return }

    /* Iterate over the slice and send the pending messages */
    for _,v := range s.PendingMsg[msg.Id] {
      err = gob.NewEncoder(conn).Encode(v)
      if err != nil { log.Println(err) }
    }
    delete(s.PendingMsg, msg.Id)

  case DISCONNECT_CODE: // Remove user from the list of users
    delete(s.Users, msg.Id)

  default:
    log.Println("Message=", msg)
  }

}
