package main

import (
  "log"
  "fmt"
  "labs"
  "labs/rpclab"
  "labs/data"
  "net/rpc"
)

func main() {
  var avg float64
  var opt string
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
      if err != nil {
        fmt.Println(rpclab.MADD_ONE, err)
        continue
      }
      fmt.Println("Calificacion agregada exitosamente!\n")

    case "b":
      /* Get the map of students */
      names := &rpclab.NamesList{}
      err = conn.Call(rpclab.MGET_NAMES, &dummy, names)
      if err != nil {
        fmt.Println(rpclab.MGET_NAMES, err)
        continue
      }

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
      if err != nil {
        fmt.Println(rpclab.MAVG_ONE, err)
        continue
      }

      /* Show the average */
      fmt.Printf("El promedio de %s es %.2f\n", val, avg)

    case "c":
      /* Call the method from the server to get the overall average */
      err = conn.Call(rpclab.MAVG_ALL, &dummy, &avg)
      if err != nil {
        fmt.Println(rpclab.MAVG_ALL, err)
        continue
      }

      /* Show the average */
      fmt.Printf("El promedio de todos los alumnos es %.2f\n", avg)

    case "d":
      // TODO: Promedio por materia
      /* Get the slice of subjects */
      list := &rpclab.SubjectsList{}
      err = conn.Call(rpclab.MGET_SUBS, dummy, &list)
      if err != nil {
        fmt.Println(rpclab.MGET_SUBS, err)
        continue
      }

      /* Display all the subjects */
      for i,s := range list.Value {
        fmt.Println(i, ":", s)
      }
      fmt.Print("Escoja el numero de la materia: ")
      fmt.Scanf("%d", &id)

      if id < 0 || id >= len(list.Value) {
        fmt.Println("Numero de materia invalido")
        continue
      }

      /* Get the average from the choosen subject */
      err = conn.Call(rpclab.MAVG_SUB, list.Value[id], &avg)
      if err != nil {
        fmt.Println(rpclab.MAVG_SUB, err)
        continue
      }
      fmt.Printf("El promedio de la materia %s es %.2f\n", list.Value[id], avg)

    case "x":
      fmt.Println("Terminando el programa")
      break

    default:
      fmt.Println("Opcion incorrecta")
    }
  }
}

func readGrade() * data.Grade {
  var grade float64

  fmt.Print("Ingrese el nombre del alumno: ")
  name, err := labs.ReadLine()
  if err != nil { log.Fatal("Read input:", err) }

  fmt.Print("Ingrese el nombre de la materia: ")
  sub, err := labs.ReadLine()
  if err != nil { log.Fatal("Read input:", err) }

  fmt.Print("Ingrese la calificacion: ")
  fmt.Scanf("%f", &grade)

  return data.NewGrade(name, sub, grade)
}

func menu() {
  fmt.Println("\na) Agregar la calificación de un alumno por materia.")
  fmt.Println("b) Obtener el promedio del alumno.")
  fmt.Println("c) Obtener el promedio de todos los alumnos.")
  fmt.Println("d) Obtener el promedio por materia.")
  fmt.Println("x) Salir.\n")
  fmt.Print("Escoja una opcion: ")
}
