package labs

import (
  "log"
  "net"
  "encoding/gob"
)

type Server struct {
  ln net.Listener
  users map[int]string
}

type server interface {
  Init()
  Run()
  HandleClient(conn net.Conn)
}

/*
func HandleClient(conn net.Conn, users map[int]string) {
  var msg Message

  err := gob.NewDecoder(conn).Decode(&msg)

  if err != nil {
    log.Println(err)
    return
  }

  switch msg.Code {
  case REGISTER_CODE:
    users[msg.Id] = msg.Data
    log.Println("User registered")
  default:
    log.Println("Message=", msg)
  }

  conn.Close()
}
*/

func (s * Server) Init() {
  s.users = make(map[int]string)
  ln, err := net.Listen(PROTOCOL, PortString)
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
    /* accept connection on port */
    conn, err := s.ln.Accept()
    if err != nil { log.Println(err, "ln.Accept()") }
    go s.HandleClient(conn)
    log.Println(s.users)
  }
}

func (s * Server) HandleClient(conn net.Conn) {
  var msg Message

  err := gob.NewDecoder(conn).Decode(&msg)

  if err != nil {
    log.Println(err)
    return
  }

  switch msg.Code {
  case REGISTER_CODE:
    s.users[msg.Id] = msg.Data
    log.Println("User registered")
  default:
    log.Println("Message=", msg)
  }

  conn.Close()
}
