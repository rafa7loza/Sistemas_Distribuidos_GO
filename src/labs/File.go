package labs

import (
  "os"
)

type File struct {
  Name string
  Content []byte
}

/* Export functions */
func AppendToFile(filename string, msg []byte) error {
  fd, err := os.OpenFile(
    filename,
    os.O_APPEND|os.O_CREATE|os.O_WRONLY,
    0644)

  if err != nil { return nil }
  defer fd.Close()

  defer fd.Close()
  if _, err := fd.Write(msg); err != nil { return err }
  return nil
}
