package labs

import (
  "log"
  "net"
  "encoding/gob"
)

type Server struct {
  ln net.Listener
  Users map[int]string
}

type server interface {
  Init()
  Run()
  HandleClient(conn net.Conn)
}

/* Server methods */
func (s * Server) Init() {
  s.Users = make(map[int]string)
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

  err := gob.NewDecoder(conn).Decode(&msg)
  if err != nil {
    log.Println(err)
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
      log.Println(err)
      return
    }
  default:
    log.Println("Message=", msg)
  }

}
