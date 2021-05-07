package main

import (
  "log"
  "fmt"
  "labs"
  "labs/rpclab"
  "net/rpc"
)

func main() {
  opt := ""
  conn, err := rpc.Dial("tcp", labs.LOCAL + ":" + labs.PORT)
  if err != nil { log.Fatal("Fallo al conectarse al servidor:", err) }
  defer conn.Close()

  for ; opt != "x" ; {
    menu()
    fmt.Scanln(&opt)

    switch opt {
    case "a":
      // TODO: Agregar calificacion
      grade := readGrade()
      var tmp int
      err = conn.Call("Server.AddGrade", grade, &tmp)
      if err != nil { log.Fatal("Server.AddGrade", err) }
      log.Println("TODO: Agregar calificacion")

    case "b":
      // TODO: Obtener promedio
      log.Println("TODO: Obtener promedio")

    case "c":
      // TODO: Promedio de todos los alumnos
      log.Println("TODO: Promedio de todos los alumnos")

    case "d":
      // TODO: Promedio por materia
      log.Println("TODO: Promedio por materia")

    case "x":
      log.Println("Terminando el programa")
      break

    default:
      log.Println("Opcion incorrecta")
    }
  }
}

func readGrade() * rpclab.Grade {
  var grade float64

  fmt.Print("Ingrese el nombre del alumno: ")
  name, err := labs.ReadLine()
  if err != nil { log.Fatal("Read input:", err) }

  fmt.Print("Ingrese el nombre de la materia: ")
  sub, err := labs.ReadLine()
  if err != nil { log.Fatal("Read input:", err) }

  fmt.Print("Ingrese la calificacion: ")
  fmt.Scanf("%f", &grade)

  return rpclab.NewGrade(name, sub, grade)
}

func menu() {
  fmt.Println("a) Agregar la calificación de un alumno por materia.")
  fmt.Println("b) Obtener el promedio del alumno.")
  fmt.Println("c) Obtener el promedio de todos los alumnos.")
  fmt.Println("d) Obtener el promedio por materia.")
  fmt.Println("x) Salir.\n")
  fmt.Print("Escoja una opcion: ")
}
