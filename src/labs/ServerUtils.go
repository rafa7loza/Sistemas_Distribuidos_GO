package labs

import (
  "io"
  "log"
  "net"
  "fmt"
  "time"
  "encoding/json"
  "encoding/gob"
)

const (
  SENDFILE = 1
  SENDMSG = 2
)

type Server struct {
  ln net.Listener
  Users map[int]string
  PendingMsg map[int][]Message
  Logs []string
}

type server interface {
  Init()
  Run()
  HandleClient(conn net.Conn)
}

/* Internal functions */
func formatLog(source, dest, extra string, action int) string {
  var formatStr string
  dt := time.Now()

  if action == SENDFILE {
    formatStr = "%s: %s sent the file '%s' to %s"
  } else if action == SENDMSG {
    formatStr = "%s: %s sent the message '%s' to %s"
  } else {
    return ""
  }

  return fmt.Sprintf(
    formatStr,
    dt.Format("01-02-2006 15:04:05"),
    source,
    extra,
    dest)
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
  var file File

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

  case GETUSERS_CODE: // Send the map of users
    msg = *NewMessage(GETUSERS_CODE, -1, -1, s.Users)
    err = gob.NewEncoder(conn).Encode(msg)
    if err != nil {
      if err != io.EOF { log.Println(err) }
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

    /* Format and store the log */
    log := formatLog(
      s.Users[msg.Id],
      s.Users[msg.Dest],
      string(msg.Data),
      SENDMSG)
    s.Logs = append(s.Logs, log)

    /* Create a message with the text and add it to the waiting list */
    msg = *NewMessage(RECMESSAGE_CODE, msg.Id, msg.Dest, []byte(text))
    s.PendingMsg[msg.Dest] = append(s.PendingMsg[msg.Dest], msg)

    /* Return the message */
    msg = *NewMessage(OK_CODE, -1, -1, []byte{})
    gob.NewEncoder(conn).Encode(msg)

  case SENDFILE_CODE:
    /* Decode the File struct */
    err := json.Unmarshal(msg.Data, &file)
    if err != nil {
      log.Println(err)
      return
    }

    /* Format and store the log */
    log := formatLog(
      s.Users[msg.Id],
      s.Users[msg.Dest],
      string(file.Name),
      SENDFILE)
    s.Logs = append(s.Logs, log)

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

func (s * Server) PrintLogs() {
  for _,log := range s.Logs {
    fmt.Println(log)
  }
  fmt.Println()
}
