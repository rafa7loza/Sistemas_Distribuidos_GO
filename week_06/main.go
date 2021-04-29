package main

import (
  "fmt"
  "sort"
  "os"
  "log"
);

type stringSlice []string;
func (arr stringSlice) Len() int { return len(arr) }
func (arr stringSlice) Swap(i, j int) { arr[i], arr[j] = arr[j], arr[i] }
func (arr stringSlice) Less(i, j int) bool { return arr[i] < arr[j] }

func writeToFile(filename string, arr stringSlice) error {
  fp, err := os.OpenFile(filename, os.O_RDWR | os.O_CREATE, 0755);
  if err != nil { return err; }

  for _,v := range arr {
    _, err := fp.WriteString(v+"\n");
    if err != nil { return err; }
  }

  if err := fp.Close(); err != nil { return err; }

  return nil;
}

func main() {
  var n int64;

  fmt.Print("Ingrese en número de cadenas a leer: ");
  fmt.Scanf("%d", &n);
  asc := make([]string, n);
  desc := make([]string, n);

  for i,_ := range asc {
    fmt.Scanf("%s", &asc[i]);
    desc[i] = asc[i];
  }

  sort.Sort(stringSlice(asc));
  sort.Sort(sort.Reverse(stringSlice(desc)));


  if err := writeToFile("asecendente.txt", asc); err != nil {
    log.Fatal(err);
  }

  if err := writeToFile("descendente.txt", desc); err != nil {
    log.Fatal(err);
  }

  fmt.Println("Las cadenas se han escrito exitosamente en ambos archivos");

}
