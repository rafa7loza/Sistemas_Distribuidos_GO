package utils

import (
  "time"
  "fmt"
)

type Message struct {
  From    string  `json:"From"`
  Content string  `json:"Content"`
  Date    string  `json:"Date"`
}

func NewMessage(user, content string) * Message {
  currentTime := time.Now()
  date := currentTime.Format("2006-01-02 15:04:05")
  return &Message{user, content, date}
}

func (msg * Message) ToString() string {
  return fmt.Sprintf(
    "%s %s: %s",
    msg.Date,
    msg.From,
    msg.Content,
  )
}
