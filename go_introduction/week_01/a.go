package main

import "fmt"

func main(){
  var side float64;
  fmt.Println("\tArea de un rectangulo\n");
  fmt.Print("Lado: ");
  fmt.Scanf("%f", &side);
  area := side * side;
  fmt.Printf("Area = %.2f\n", area);
}
