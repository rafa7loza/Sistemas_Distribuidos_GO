package httplab

import (
  "errors"
)

type DataStudents struct {
  students map[string]Student
  cnt int
}

/* Constructor */
func NewDataStudents() * DataStudents {
  return &DataStudents{make(map[string]Student), 0}
}

/* Private methods */
func (data * DataStudents) addStudent(name string) {
  if data.hasStudent(name) { return }
  data.students[name] = *NewStudent(data.cnt)
  data.cnt++
}

func (data * DataStudents) hasStudent(name string) bool {
  _, ok := data.students[name]
  return ok
}

/* Public methods */
func (data * DataStudents) AddGrade(grade *Grade) error {
  data.addStudent(grade.NameStudent)
  student, _ := data.students[grade.NameStudent]
  return student.AddGrade(grade.Subject, grade.Grade)
}

func (data * DataStudents) GetAvgOne(name string, avg * float64) error {
  if !data.hasStudent(name) { return errors.New("Student is not stored") }

  *avg = 0.0
  for _,v := range data.students[name].subjects {
    *avg += v
  }
  *avg /= float64(len(data.students[name].subjects))
  return nil
}

func (data * DataStudents) GetSubjects() []string {
  arr := make([]string,0)
  tmp := make(map[string]bool)
  for _,student := range data.students {
    for sub,_ := range student.subjects {
      tmp[sub] = true
    }
  }
  for v,_ := range tmp {
    arr = append(arr, v)
  }
  return arr
}

func (data * DataStudents) GetStudents() []string {
  arr := make([]string,0)
  for name,_ := range data.students {
    arr = append(arr, name)
  }

  return arr
}
