package rpclab

import (
  "errors"
)

const (
  MADD_ONE = "Server.AddGrade"
  MAVG_ONE = "Server.GetAvgOne"
  MAVG_ALL = "Server.GetAvgAll"
  MAVG_SUB = "Server.GetAvgSub"
  MGET_NAMES = "Server.GetStudents"
  MGET_SUBS = "Server.GetSubjects"
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
  var individualAvg float64

  for _,student := range s.students {
    individualAvg = 0.0
    for _,v := range student.subjects {
      individualAvg += v
    }
    individualAvg /= float64(len(student.subjects))
    *avg += individualAvg
  }

  *avg /= float64(len(s.students))
  return nil
}

func (s * Server) GetAvgSub(subject string, avg * float64) error {
  *avg = 0.0
  var n int

  for _,student := range s.students {
    grade,ok := student.subjects[subject]
    if !ok { continue }
    *avg += grade
    n++
  }

  if n == 0 {
    return errors.New("This subject is not stored")
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

func (s * Server) GetSubjects(arg1 int, subjects * SubjectsList) error {
  subjects.Value = make([]string, 0)
  tmp := make(map[string]bool)
  for _,student := range s.students {
    for sub,_ := range student.subjects {
      if ok := tmp[sub]; ok { continue }
      subjects.Value = append(subjects.Value, sub)
      tmp[sub] = true
    }
  }
  return nil
}
