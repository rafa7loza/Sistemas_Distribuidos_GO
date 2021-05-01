package labs

const (
  PortString = ":1234"
  PROTOCOL = "tcp"
  REGISTER_CODE = 1
  WAIT_TIME_MS = 500
)

type Message struct {
  Code, Id, Dest int
  Data string
}

/* External function */

func NewMessage(code int, pid int, data string) * Message {
  return &Message{code, pid, -1, data}
}
