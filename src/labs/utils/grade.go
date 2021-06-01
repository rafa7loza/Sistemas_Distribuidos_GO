package utils

type Grade struct {
  NameStudent string  `json:NameStudent`
  Subject string      `json:Subject`
  Grade float64       `json:Grade`
}

func NewGrade(name, sub string, grade float64) * Grade {
  return &Grade{name, sub, grade}
}
