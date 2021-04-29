package main

import (
  "fmt"
  "time"
)

var showLogs bool = false;

type Process interface {
  Run();
  Stop();
  isRunning() bool;
  GetId() int64;
}

type P struct {
  id int64;
  stoper chan bool;
  running bool;
}

type Processes struct {
  arr []Process;
}

func (p P) Run() {
  i := int64(0);
  p.running = true;

  for {
    select {
    case <-p.stoper:
      p.running = false;
      for {
        _, ok := <- p.stoper;
        if !ok { break; }
      }
      close(p.stoper);
      return ;
    default:
      if showLogs {
        fmt.Printf("{id=#%d, i=%d}\n", p.id, i);
      }
      i++;
      time.Sleep(time.Millisecond * 500);
    }
  }

}

func (p P) Stop() {
  p.stoper <- true;
}

func (p P) isRunning() bool {
  return p.running;
}

func (p P) GetId() int64 {
  return p.id;
}

func menu() {
  fmt.Println("\n\tAdministrador de procesos");
  fmt.Println("a) Agregar proceso.");
  fmt.Println("b) Eliminar proceso.");
  fmt.Println("c) Mostrar procesos (presione ENTER para dejar de mostrar).");
  fmt.Println("x) Salir.");
  fmt.Print("\nEscoja una opción: ");
}

func main() {
  var opt string;
  var p Process;
  var _pid int64;

  ps := Processes{};
  autoIncId := int64(0);

  for {
    menu();
    fmt.Scan(&opt);
    if opt == "x" { break; }

    switch opt {
    case "a":
      p = P{ autoIncId, make(chan bool), false};
      go p.Run();
      ps.arr = append(ps.arr, p);
      autoIncId++;
    case "b":
      fmt.Print("Ingrese el ID del proceso que desea terminar: ");
      fmt.Scanf("%d", &_pid);
      if _pid < 0 || _pid >= int64(len(ps.arr)) {
        fmt.Println("ID inválido!")
      } else {
        ps.arr[_pid].Stop();
        fmt.Printf("Se ha terminado el proceso %d\n", ps.arr[_pid].GetId());
      }
    case "c":
      showLogs = true;
      fmt.Scanf("%s", &opt);
      showLogs = false;
    default:
      fmt.Println("Opcion inválida");
    }
  }

}
