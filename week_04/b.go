package main

import "fmt"

func fib(x int64) int64 {
  if x <= 1 { return 1; } else { return fib(x-1) + fib(x-2); }
}
func main() {
  var n int64;
  fmt.Print("Ingrese hasta que numero de la secuencia quiere llegar: ");
  fmt.Scanf("%d",&n);
  fmt.Printf("Fib(%d) = %d\n", n, fib(n));
}
