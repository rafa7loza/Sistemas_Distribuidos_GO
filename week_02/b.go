package main

import "fmt"

/*
 * Calculate the euler constant
 */
func main() {
  var _e float64;
  N := int64(20);
  arr := make([]int64, N);

  for i,_ := range arr {
    if i <= 1 {
      arr[i] = int64(1);
    } else {
      arr[i] = arr[i-1] * int64(i);
    }
  }

  for _,v := range arr {
    _e += float64(float64(1.0) / float64(v));
  }

  fmt.Println("e =", _e);
}
