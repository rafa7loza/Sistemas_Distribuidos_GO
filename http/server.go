package main

import (
  "log"
  "net/http"
  "labs"
  "fmt"
)

func root(res http.ResponseWriter, req *http.Request) {
  res.Header().Set(
    "Content-Type",
    "text/html",
  )

  content, err := labs.ReadFileContent("www/index.html")
  if err != nil { log.Fatal("Read the file content") }

	fmt.Fprintf(
		res,
    content,
	)

}

func main() {
  http.HandleFunc("/", root)
  log.Println("Starting the server")
  address := labs.LOCAL + ":" + labs.PORT
  http.ListenAndServe(address, nil)
}
