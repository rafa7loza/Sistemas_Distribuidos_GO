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

  content, err := labs.ReadFileContent("index.html")
  if err != nil { log.Fatal("Read the file content") }

	fmt.Fprintf(
		res,
    content,
	)

}

func getScript(res http.ResponseWriter, req *http.Request) {
  res.Header().Set(
    "Content-Type",
    "text/javascript",
  )

  content, err := labs.ReadFileContent("js/ui.js")
  if err != nil { log.Fatal("Read the file content") }

	fmt.Fprintf(
		res,
    content,
	)
}

func calificacion(res http.ResponseWriter, req *http.Request) {
  log.Println(req.Method)
  // switch req.Method {
	// case "POST":
	// 	if err := req.ParseForm(); err != nil {
	// 		fmt.Fprintf(res, "ParseForm() error %v", err)
	// 		return
	// 	}
  //
	// 	log.Println(req.PostForm)
	// }
}

func main() {
  handler := http.NewServeMux()
  handler.HandleFunc("/", root)
  handler.HandleFunc("/calificacion", calificacion)
  handler.HandleFunc("/js/ui.js", getScript)

  log.Println("Starting the server")
  address := ":" + labs.PORT
  http.ListenAndServe(address, handler)
}
