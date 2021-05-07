package labs

import (
  "bufio"
  "os"
  "strings"
)

func ReadLine() (string, error) {
  reader := bufio.NewReader(os.Stdin)
  data, err := reader.ReadString('\n')
  if err != nil { return "(NULL)", err }
  data = strings.Replace(data, "\n", "", -1)
  return data, nil
}
