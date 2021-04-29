package main

import "fmt"

func main(){
  var h, b float64;
  fmt.Println("\tArea de un triángulo\n");
  fmt.Print("Base: ");
  fmt.Scanf("%f", &b);
  fmt.Print("Altura: ");
  fmt.Scanf("%f", &h);
  area := (b*h)/2;
  fmt.Printf("Area = %.2f\n", area);
}
