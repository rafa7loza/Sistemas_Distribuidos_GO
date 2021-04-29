package main

import "fmt"

/*
 * Read a slice from standar input and calculate the sum of
 * all the elements of this slice
 */
func main() {
  var n, x int32;
  var _sum int64;
  var arr []int32;
  fmt.Scanf("%d", &n);

  for i := int32(0); i < n; i++ {
    fmt.Scanf("%d", &x);
    arr = append(arr, x);
  }

  for _,v := range arr {
    _sum += int64(v);
  }

  fmt.Println(_sum);
}
