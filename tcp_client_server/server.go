package main

import(
  "net"
  "fmt"
  "bufio"
  "log"
  "mypkg"
)

func main() {
  fmt.Println("Starting the server...")

  /* Listen to the TCP port */
  ln, err := net.Listen("tcp", mypkg.PortString)
  if err != nil { log.Fatal(err) }

  /* Create the slice of processes and run them */
  ps := mypkg.NewProcesses(4)
  go ps.Run()

  for {
    /* accept connection on port */
    conn, err := ln.Accept()
    if err != nil { log.Fatal(err) }
    go handleClient(conn, ps)
  }
}

func handleClient(conn net.Conn, ps * mypkg.Processes) {
  var newProc mypkg.P

  /* Listen for the client, process the package reading all the bytes
     until the delimiter */
  message, err := bufio.NewReader(conn).ReadBytes(mypkg.DEL)
  if err != nil { log.Fatal(err) }

  /* Decode the package to a byte array */
  pkg, err := mypkg.DecodePackage(message);
  if err != nil { log.Fatal(err) }

  if pkg.Code == mypkg.REQUEST_CODE {
    log.Println("A client is requesting a package")
    err := ps.GiveProcess(&newProc)

    if err != nil {
      pkg.Code = mypkg.NOPROC_CODE
      log.Println(err)
    } else {
      /* Delete the given process from the array */
      copy(ps.Arr[0:], ps.Arr[1:])
      ps.Arr = ps.Arr[:len(ps.Arr)-1]

      /* Update the package data */
      pkg.Code = mypkg.GETPROC_CODE
      pkg.Id = newProc.Id
      pkg.Cnt = newProc.Cnt
    }
  } else if pkg.Code == mypkg.ENDP_CODE {
    /* The client finished processing the package */
    log.Println("Received a package from a client")
    proc := mypkg.NewProcess(pkg.Id, pkg.Cnt)
    ps.Arr = append(ps.Arr, *proc)
    return
  } else {
    log.Println("Package not recognized")
  }

  /* Reply to the client */
  n, err := conn.Write(pkg.Encode())
  log.Printf("%d bytes were sent.\n", n);
  if err != nil { log.Fatal(err); }

  conn.Close()
}
