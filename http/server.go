package main

import (
  "log"
  "net/http"
  "labs"
  "labs/http"
  "fmt"
  "strconv"
  "encoding/json"
)

var students * httplab.DataStudents

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

  content, err := labs.ReadFileContent("js/util.js")
  if err != nil { log.Fatal("Read the file content") }

	fmt.Fprintf(
		res,
    content,
	)
}

func getSubjects(res http.ResponseWriter, req *http.Request) {
  switch req.Method {
	case "GET":
    res.Header().Set(
      "Content-Type",
      "application/json",
    )

    mm := make(map[string][]string, 0)
    mm["subjects"] = students.GetSubjects()
    log.Println(mm)

    content, err := json.Marshal(mm)
    if err != nil {
      log.Println(err)
      return
    }

    fmt.Fprintf(
      res,
      string(content),
    )
  }
}

func getStudents(res http.ResponseWriter, req *http.Request) {
  switch req.Method {
	case "GET":
    res.Header().Set(
      "Content-Type",
      "application/json",
    )

    mm := make(map[string][]string, 0)
    mm["students"] = students.GetStudents()
    log.Println(mm)

    content, err := json.Marshal(mm)
    if err != nil {
      log.Println(err)
      return
    }

    fmt.Fprintf(
      res,
      string(content),
    )
  }
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

    /* Get the input into the correct format */
    subject = req.FormValue("studentSubject")
    fgrade, err := strconv.ParseFloat(req.FormValue("studentGrade"), 64)
    if err != nil {
      log.Println("Error converting grade input to float64")
      return
    }

    /* Add a new student to memory */
    students.AddGrade(httplab.NewGrade(
      name,
      subject,
      fgrade))

    htmlContent, err := labs.ReadFileContent("response.html")
    if err != nil {
      log.Println(err)
      return
    }

    msg := fmt.Sprintf("Se ha agregado el alumno %s correctamente\n", name)

    res.Header().Set(
			"Content-Type",
			"text/html",
		)

		fmt.Fprintf(
			res,
			htmlContent,
			msg,
		)
	}
}

func main() {
  students = httplab.NewDataStudents()
  handler := http.NewServeMux()
  handler.HandleFunc("/", root)
  handler.HandleFunc("/calificacion", calificacion)
  handler.HandleFunc("/js/util.js", getScript)
  handler.HandleFunc("/data/subjects.json", getSubjects)
  handler.HandleFunc("/data/students.json", getStudents)

  log.Println("Starting the server")
  address := ":" + labs.PORT
  http.ListenAndServe(address, handler)
}
