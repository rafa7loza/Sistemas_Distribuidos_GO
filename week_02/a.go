package main

import "fmt"

func main() {
  var d,m int;
  signs := [12]string{"Capricornio","Acurario","Piscis","Aries","Tauro","Géminis","Cáncer","Leo","Virgo","Libra","Escorpio","Sagitario"}
  fmt.Scanf("%d", &d);
  fmt.Scanf("%d", &m);

  switch m {
    case 1:
      if d < 21 {
        fmt.Println(signs[0]);
      } else {
        fmt.Println(signs[1]);
      }
    case 2:
      if d < 20 {
        fmt.Println(signs[1]);
      } else {
        fmt.Println(signs[2]);
      }
    case 3:
      if d < 21 {
        fmt.Println(signs[2]);
      } else  {
        fmt.Println(signs[3]);
      }
    case 4:
      if d < 21 {
        fmt.Println(signs[3]);
      } else {
        fmt.Println(signs[4]);
      }
    case 5:
      if d < 21 {
        fmt.Println(signs[4]);
      } else {
        fmt.Println(signs[5]);
      }
    case 6:
      if d < 22 {
        fmt.Println(signs[5]);
      } else {
        fmt.Println(signs[6]);
      }
    case 7:
      if d < 24 {
        fmt.Println(signs[6]);
      } else {
        fmt.Println(signs[7]);
      }
    case 8:
      if d < 24 {
        fmt.Println(signs[7]);
      } else {
        fmt.Println(signs[8]);
      }
    case 9:
      if d < 24 {
        fmt.Println(signs[8]);
      } else {
        fmt.Println(signs[9]);
      }
    case 10:
      if d < 24 {
        fmt.Println(signs[9]);
      } else {
        fmt.Println(signs[10]);
      }
    case 11:
      if d < 23 {
        fmt.Println(signs[10]);
      } else {
        fmt.Println(signs[11]);
      }
    case 12:
      if d < 22 {
        fmt.Println(signs[11]);
      } else {
        fmt.Println(signs[0]);
      }
    default:
      fmt.Println("Entrada inválida");
  }
}
