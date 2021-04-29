package main

import "fmt"

const MinInt64 = (^int64(0)) << 63;

func greater(args ...int64) int64 {
  max := MinInt64;
  for _,v := range args {
    if max < v {
      max = v;
    }
  }
  return max;
}

func main() {
  values := []int64{1,3,5,3,12,5,2,42,6,-1};
  fmt.Println(greater(values...));
}
