package utils

import (
  "errors"
)

type Chat struct {
  Msgs  []Message       `json:"Messages"`
  Users map[string]bool `json:"Users"`
}

// Public functions

/* Add a user to the map, if the user is in the map and his value it's true
 * then return an error */
func (ch * Chat) AddUser(username string) error {
  if isActive := ch.Users[username]; isActive {
    return errors.New("User is currently active")
  }

  ch.Users[username] = true
  return nil
}

func (ch * Chat) AddMessage(username, content string) error {
  msg := NewMessage(username, content)
  if msg == nil {
    return errors.New("Failed to create the message")
  }

  ch.Msgs = append(ch.Msgs, *msg)
  return nil
}

func (ch * Chat) GetMessages() []string {
  var ret []string

  for _,msg := range ch.Msgs {
    ret = append(ret, msg.ToString())
  }

  return ret
}
