package main

import(
  "net"
  "fmt"
  "bufio"
  "mypkg"
  "log"
  "os"
  "os/signal"
)

func getProcess() * mypkg.Pkg {
  /* Do the first connection to receive a process */
  conn, err := net.Dial("tcp", mypkg.PortString)
  if err != nil { log.Fatal(err) }

  /* Create a package indicating that the client wants to receive
     a process, then send it to the server */
  pkg := mypkg.NewPackage(mypkg.REQUEST_CODE)
  conn.Write(pkg.Encode())

  /* Listen for reply */
  message, err := bufio.NewReader(conn).ReadBytes(mypkg.DEL)
  if err != nil { log.Fatal(err) }

  /* Decode the package */
  newPkg, err := mypkg.DecodePackage(message);
  if err != nil { log.Fatal(err) }

  switch newPkg.Code {
  case mypkg.GETPROC_CODE:
    fmt.Println("Package received successfully")
    break
  case mypkg.NOPROC_CODE:
    fmt.Println("There are not process availabe")
    break
  default:
    fmt.Println("Unknown code")
    break
  }

  /* Terminate the connection and return the package */
  conn.Close()

  return newPkg
}

func returnProcess(proc * mypkg.P) bool {
  /* Do the second connection to give the process back */
  conn, err := net.Dial("tcp", mypkg.PortString)
  if err != nil { log.Fatal(err) }

  /* Update the code */
  pkg := mypkg.NewPackage(mypkg.ENDP_CODE)
  pkg.Cnt = proc.Cnt
  pkg.Id = proc.Id

  /* Encode and send the package to the server */
  conn.Write(pkg.Encode())

  /* Close connection */
  conn.Close()
  return true
}

func main() {

  /* First connection */
  pkg := getProcess()
  if pkg == nil { log.Fatal("Failed to receive a process") }

  /* Convert package to a Process */
  proc := mypkg.NewProcess(pkg.Id, pkg.Cnt)
  go proc.Run()

  signalChannel := make(chan os.Signal, 1)
	signal.Notify(signalChannel, os.Interrupt)
	// Block until a signal is received.
	s := <-signalChannel
  fmt.Println("Ending the programm", s)

  /*Stop the current process */
  proc.Stop()

  /* Send the package back to the server */
  ok := returnProcess(proc)
  if ok {
    log.Println("Process ended successfully!")
  }

}
