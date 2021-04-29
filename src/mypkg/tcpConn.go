package mypkg

import (
  "bytes"
  "encoding/binary"
  "log"
  "fmt"
  "time"
  "errors"
)

type P struct {
  Id, Cnt int64
  Stoper chan bool
}

type Pkg struct {
  Code, Id, Cnt int64
}

type Package interface {
  Encode();
  Run();
}

type Processes struct {
  Arr []P
}

type Process interface {
  Init(n int64)
  Run()
  Stop()
  GiveProcess() error
}

const (
  REQUEST_CODE = int64(7)
  GETPROC_CODE = int64(8)
  ENDP_CODE = int64(9)
  NOPROC_CODE = int64(10)
  CLIENT_CODE = int64(11)
  Port = 1234;
  PortString = ":1234";
  DEL = '\n'
  WAIT_TIME_MS = 500
)


func NewProcesses(n int64) *Processes {
  var ps Processes
  for i := int64(0); i<n; i++ {
    ps.Arr = append(ps.Arr, *NewProcess(i, 0))
  }
  return &ps
}

func (ps Processes) GiveProcess(p * P) error {
  n := len(ps.Arr)

  if n == 0 {
    return errors.New("No process available")
  }

  *p = ps.Arr[0]
  return nil
}

func (ps * Processes) Run() {
  for {
    for i,v := range ps.Arr {
      fmt.Printf("%d : %d\n", v.Id, v.Cnt)
      ps.Arr[i].Cnt++
    }
    time.Sleep(time.Millisecond * WAIT_TIME_MS);
    fmt.Println("----------------------")
  }
}

func NewProcess(id, cnt int64)  *P  { return &P{id, cnt, make(chan bool, 1)} }

func NewPackage(code int64) * Pkg {
  return &Pkg{
    code,
    0,
    0}
}

func DecodePackage(b []byte) (*Pkg, error) {
  scanner := bytes.NewReader(b)
  var pkg Pkg
  if err := binary.Read(scanner, binary.LittleEndian, &pkg); err != nil {
    return nil, err
  }

  return &pkg, nil
}

func (pkg * Pkg) Encode() []byte {
  buffer := new(bytes.Buffer)
  err := binary.Write(buffer, binary.LittleEndian, pkg)
  if err != nil { log.Fatal(err) }
  b := buffer.Bytes()
  b = append(b, byte(DEL))
  return b
}

func (proc * P) Run() {
  for {
    select {
    case <-proc.Stoper:
      for {
        _, ok := <- proc.Stoper
        if !ok { break }
      }
      close(proc.Stoper)
      return
    default:
      fmt.Printf("%d : %d\n", proc.Id, proc.Cnt)
      proc.Cnt++
      time.Sleep(time.Millisecond * WAIT_TIME_MS);
      fmt.Println("----------------------")
    }
  }
}

func (proc * P) Stop() {
  proc.Stoper <- true
}
