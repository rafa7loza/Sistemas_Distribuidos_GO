package rpclab

import (
  "log"
)

type Server struct {
  students map[string]Student
  cnt int
}

func NewServer() * Server {
  return &Server{make(map[string]Student), 0}
}

/* Public methods */
func (s * Server) addStudent(name string) {
  if s.hasStudent(name) { return }
  s.students[name] = *NewStudent(s.cnt)
  s.cnt++
}

/* Private methods */
func (s * Server) hasStudent(name string) bool {
  _, ok := s.students[name]
  return ok
}


/* RPC methods */
func (s * Server) AddGrade(grade *Grade, reply * int) error {

  log.Println("Calling the method AddGrade")
  s.addStudent(grade.NameStudent)
  log.Println(s)

  return nil
}
