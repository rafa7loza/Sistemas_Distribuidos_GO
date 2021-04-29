package main

import "fmt"

func intercambiar(a *int64, b *int64) {
  tmp := *a;
  *a = *b;
  *b = tmp;
}


func main() {
  var a,b int64;
  fmt.Scanf("%d %d", &a, &b);
  intercambiar(&a,&b);
  fmt.Printf("%d %d\n", a, b);
}
