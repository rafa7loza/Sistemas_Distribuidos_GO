package main

import "fmt"

func generadorImpares() func() uint64 {
  it := uint64(1);
  return func() uint64 {
    var nxt = it;
    it += 2;
    return nxt;
  }
}
func main() {
  nxt := generadorImpares();
  for i := 0; i<10; i++ {
    fmt.Println(nxt());
  }
}
