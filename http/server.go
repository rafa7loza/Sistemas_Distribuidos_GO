package main

import (
  "log"
  "net/http"
  "labs"
  "labs/http"
  "fmt"
  "strconv"
)

var students * httplab.Data

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
  switch req.Method {
	case "POST":
		if err := req.ParseForm(); err != nil {
			fmt.Fprintf(res, "ParseForm() error %v", err)
			return
		}

		log.Println(req.PostForm)
    var name, subject string
    var fgrade float64

    if _,ok := req.PostForm["studentExists"]; !ok {
      log.Println("Add student to the list")
      name = req.FormValue("studentName")
    } else {
      log.Println("Student already added")
      name = req.FormValue("names")
    }

    subject = req.FormValue("studentSubject")
    fgrade, err := strconv.ParseFloat(req.FormValue("studentGrade"), 64)
    if err != nil {
      log.Println("Error converting grade input to float64")
      return
    }


    students.AddGrade(httplab.NewGrade(
      name,
      subject,
      fgrade))

    log.Println(students)
	}
}

func main() {
  students = httplab.NewData()
  handler := http.NewServeMux()
  handler.HandleFunc("/", root)
  handler.HandleFunc("/calificacion", calificacion)
  handler.HandleFunc("/js/ui.js", getScript)

  log.Println("Starting the server")
  address := ":" + labs.PORT
  http.ListenAndServe(address, handler)
}
