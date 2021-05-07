package rpclab

type Student struct {
  id int
  subjects map[string]float64
}

/* Constructor */
func NewStudent(id int) * Student {
  return &Student{id, make(map[string]float64) }
}
