package rpclab

import (
  "log"
  "errors"
)

const (
  MADD_ONE = "Server.AddGrade"
  MAVG_ONE = "Server.GetAvgOne"
  MAVG_ALL = "Server.GetAvgAll"
  MAVG_SUB = "Server.GetAvgSub"
  MGET_NAMES = "Server.GetStudents"
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
  s.addStudent(grade.NameStudent)
  student, _ := s.students[grade.NameStudent]
  return student.AddGrade(grade.Subject, grade.Grade)
}

func (s * Server) GetAvgOne(name string, avg * float64) error {
  log.Println(s)

  if !s.hasStudent(name) { return errors.New("Student is not stored") }

  *avg = 0.0
  for _,v := range s.students[name].subjects {
    *avg += v
  }
  *avg /= float64(len(s.students[name].subjects))
  return nil
}

func (s * Server) GetAvgAll(arg1 * int, avg * float64) error {
  *avg = 0.0
  var n int

  for _,student := range s.students {
    n += len(student.subjects)
    for _,v := range student.subjects {
      *avg += v
    }
  }

  *avg /= float64(n)
  return nil
}

func (s * Server) GetStudents(arg1 * int, names * NamesList) error {
  names.Value = make(map[int]string, 0)
  for k,v := range s.students {
    if _, ok := names.Value[v.Id]; ok {
      continue
    }
    names.Value[v.Id] = k
  }
  return nil
}
