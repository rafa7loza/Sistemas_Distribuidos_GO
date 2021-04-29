package main

import (
  "fmt"
)

func main(){
  var f float64;
  fmt.Println("\tConvertir grados Fahrenheit a Celcius\n");
  fmt.Print("Fahrenheit: ");
  fmt.Scanf("%f", &f);
  celcius := ((f - 32) * 5) / 9;
  fmt.Printf("Celcius = %.2f\n", celcius);
}
