package labs

import (
  "os"
  "errors"
)

type File struct {
  Name string
  Content []byte
}

var OPEN_ERR = errors.New("Failed to open the file")
var STAT_ERR = errors.New("Failed to stat the file")
var READ_ERR = errors.New("Failed to read the file")

/* Export functions */
func AppendToFile(filename string, msg []byte) error {
  fd, err := os.OpenFile(
    filename,
    os.O_APPEND|os.O_CREATE|os.O_WRONLY,
    0644)

  if err != nil { return OPEN_ERR }
  defer fd.Close()

  if _, err := fd.Write(msg); err != nil { return err }
  return nil
}

func ReadFileContent(filename string) (string, error) {
  fd, err := os.OpenFile(
    filename,
    os.O_RDONLY,
    0644)

  if err != nil { return "", OPEN_ERR }
  defer fd.Close()

  info, err := fd.Stat()
  if err != nil  { return "", STAT_ERR }

  N := info.Size()
  bytes := make([]byte, N)
  _, err = fd.Read(bytes)
  if err != nil { return "", READ_ERR }

  return string(bytes), nil

}
