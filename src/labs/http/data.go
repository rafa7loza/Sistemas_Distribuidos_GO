package httplab

type Data struct {
  students map[string]Student
  cnt int
}

/* Constructor */
func NewData() * Data {
  return &Data{make(map[string]Student), 0}
}

/* Private methods */
func (data * Data) addStudent(name string) {
  if data.hasStudent(name) { return }
  data.students[name] = *NewStudent(data.cnt)
  data.cnt++
}

func (data * Data) hasStudent(name string) bool {
  _, ok := data.students[name]
  return ok
}

/* Public methods */
func (data * Data) AddGrade(grade *Grade) error {
  data.addStudent(grade.NameStudent)
  student, _ := data.students[grade.NameStudent]
  return student.AddGrade(grade.Subject, grade.Grade)
}
