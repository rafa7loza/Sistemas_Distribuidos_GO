package utils

import (
  "errors"
)

type Student struct {
  Id int                        `json:Id`
  Subjects map[string]float64   `json:Subjects`
}

/* Constructor */
func NewStudent(id int) * Student {
  return &Student{id, make(map[string]float64) }
}

/* Public method */
func (st * Student) HasGrade(subject string) bool {
  _, ok := st.Subjects[subject]
  return ok
}

func (st * Student) AddGrade(subject string, grade float64) error {
  if st.HasGrade(subject) {
    return errors.New("Student already has a grade for this subject")
  }

  st.Subjects[subject] = grade
  return nil
}
