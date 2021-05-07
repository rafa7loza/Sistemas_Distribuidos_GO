package main

import (
  "log"
  "fmt"
  "labs"
  "labs/rpclab"
  "net/rpc"
)

func main() {
  var avg float64
  var name, opt string
  var dummy, id int
  opt = ""

  conn, err := rpc.Dial("tcp", labs.LOCAL + ":" + labs.PORT)
  if err != nil { log.Fatal("Fallo al conectarse al servidor: ", err) }
  defer conn.Close()

  for ; opt != "x" ; {
    menu()
    fmt.Scanln(&opt)

    switch opt {
    case "a":
      grade := readGrade()
      err = conn.Call(rpclab.MADD_ONE, grade, &dummy)
      if err != nil { log.Println(rpclab.MADD_ONE, " ", err) }
      fmt.Println("Calificacion agregada exitosamente!\n")

    case "b":
      /* Get the map of students */
      names := &rpclab.NamesList{}
      conn.Call(rpclab.MGET_NAMES, &dummy, names)
      err = conn.Call(rpclab.MGET_NAMES, name, &avg)

      /* Display the list of names */
      for k,v := range names.Value {
        fmt.Println(k, ": "+ v)
      }
      fmt.Print("Escoja el numero del alumno: ")
      fmt.Scanf("%d", &id)

      val, ok := names.Value[id]
      if !ok {
        fmt.Println("Numero de alumno invalido")
        continue
      }

      /* Call the avg method */
      err = conn.Call(rpclab.MAVG_ONE, val, &avg)
      if err != nil { log.Println(rpclab.MAVG_ONE, " ", err) }

      /* Show the average */
      fmt.Printf("El promedio de %s es %.2f\n\n", val, avg)

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
