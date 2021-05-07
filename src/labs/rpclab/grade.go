package rpclab

type Grade struct {
  NameStudent string
  Subject string
  Grade float64
}

func NewGrade(name, sub string, grade float64) * Grade {
  return &Grade{name, sub, grade}
}
