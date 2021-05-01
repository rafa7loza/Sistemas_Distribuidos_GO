package labs

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
func NewMessage(code int, pid int, data interface{}) * Message {
  switch t := data.(type) {
  case string:
    return &Message{code, pid, -1, []byte(t)}
  case []byte:
    return &Message{code, pid, -1, t}
  default:
    return nil
  }
}
