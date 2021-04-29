package main

import "fmt"

/*
 * Swap the value of two int64 variables
 *
 * @param a       Pointer to the memory address of the first variable
 * @param b       Pointer to the memory address of the second variable
 * @return None
 *
 */
func intercambiar(a *int64, b *int64) {
  tmp := *a;
  *a = *b;
  *b = tmp;
}

/*
 * Implement bubble sort algorithm.
 *
 * @param arr     Slice of int64
 * @return None
 *
 */
func Burbuja(arr []int64) {
  var i,j int64;
  off := int64(1);
  n := int64(len(arr));

  for i=0; i<n-off; i++ {
    for j=0; j<n-i-off; j++ {
      if(arr[j] > arr[j+off]){
        intercambiar(&arr[j], &arr[j+1]);
      }
    }
  }
  return -1
}


func main() {
  var arr []int64;
  var n, x int64;
  fmt.Scanf("%d", &n);

  for i := int64(0); i < n; i++ {
    fmt.Scanf("%d", &x);
    arr = append(arr, x);
  }

  for i := int64(0); i<n; i+=2 {
    Burbuja(arr);
  }
  fmt.Println(arr);

}
