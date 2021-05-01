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
  SENDMESSAGE_CODE = 3
  RECMESSAGE_CODE = 4
  CHECKMSG_CODE = 5
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
    return &Message{code, pid, dst, []byte(t)}
  case []byte:
    return &Message{code, pid, dst, t}
  case map[int]string:
    b, err := json.Marshal(t)
    if err != nil {
      log.Println(err)
      return nil
    }
    return &Message{code, pid, dst, b}
  default:
    return nil
  }
}
