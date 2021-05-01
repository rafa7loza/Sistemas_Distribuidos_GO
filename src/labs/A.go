package labs

const (
  PortString = ":1234"
  WAIT_TIME_MS = 500
)

type Message struct {
  Code, Pid int
  Data string
}

/* External function */

func NewMessage(pid int, data string) * Message {
  return &Message{1, pid, data}
}
