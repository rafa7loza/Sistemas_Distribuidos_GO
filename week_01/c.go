package main

import (
  "fmt"
  "math"
)

func main(){
  var r float64;
  fmt.Println("\tArea de un círculo\n");
  fmt.Print("Radio: ");
  fmt.Scanf("%f", &r);
  area := math.Pi * math.Pow(r,2.0);
  fmt.Printf("Area = %.2f\n", area);
}
