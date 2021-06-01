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
func (dst * DataStudents) addStudent(name string) {
  if dst.hasStudent(name) { return }
  dst.students[name] = * NewStudent(dst.cnt)
  dst.cnt++
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

func (data * DataStudents) GetAvgOne(name string) (float64, error) {
  if !data.hasStudent(name) { return 0.0, errors.New("Student is not stored") }

  avg := 0.0
  for _,v := range data.students[name].Subjects {
    avg += v
  }
  avg /= float64(len(data.students[name].Subjects))
  return avg, nil
}

func (data * DataStudents) GetAvgAll() float64 {
  avg := 0.0
  var individualAvg float64

  for _,student := range data.students {
    individualAvg = 0.0
    for _,v := range student.Subjects {
      individualAvg += v
    }
    individualAvg /= float64(len(student.Subjects))
    avg += individualAvg
  }

  avg /= float64(len(data.students))
  return avg
}

func (data * DataStudents) GetAvgSub(subject string) (float64, error) {
  avg := 0.0
  var n int

  for _,student := range data.students {
    grade,ok := student.Subjects[subject]
    if !ok { continue }
    avg += grade
    n++
  }

  if n == 0 {
    return 0.0, errors.New("This subject is not stored")
  }

  avg /= float64(n)
  return avg, nil
}

func (data * DataStudents) GetSubjects() []string {
  arr := make([]string,0)
  tmp := make(map[string]bool)
  for _,student := range data.students {
    for sub,_ := range student.Subjects {
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
