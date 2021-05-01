package labs

import (
  "encoding/json"
  "log"
)

const (
  ADDRESS = ":1234"
  PROTOCOL = "tcp"
  REGISTER_CODE = 1
  GETUSERS_CODE = 2
  WAIT_TIME_MS = 500
)

type Message struct {
  Code, Id, Dest int
  Data []byte
}

/* External functions */
func NewMessage(code int, pid int, dst int, data interface{}) * Message {
  switch t := data.(type) {
  case string:
    return &Message{code, pid, -1, []byte(t)}
  case []byte:
    return &Message{code, pid, -1, t}
  case map[int]string:
    b, err := json.Marshal(t)
    if err != nil {
      log.Println(err)
      return nil
    }
    return &Message{code, pid, -1, b}
  default:
    return nil
  }
}
