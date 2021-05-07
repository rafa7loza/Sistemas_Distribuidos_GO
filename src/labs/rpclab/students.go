package rpclab

import (
  "errors"
)

type Student struct {
  Id int
  subjects map[string]float64
}

/* Constructor */
func NewStudent(id int) * Student {
  return &Student{id, make(map[string]float64) }
}

/* Public method */
func (st * Student) HasGrade(subject string) bool {
  _, ok := st.subjects[subject]
  return ok
}

func (st * Student) AddGrade(subject string, grade float64) error {
  if st.HasGrade(subject) {
    return errors.New("Student already has a grade for this subject")
  }

  st.subjects[subject] = grade
  return nil
}
