package main

import (
  "log"
  "encoding/json"
  "errors"
  "net/http"
  "labs"
  "labs/utils"
  "strings"
  "strconv"
)

var students * utils.DataStudents

const (
  UnhandledRequest = "Unhandled request."
)

func addGrade(res http.ResponseWriter, req *http.Request) {
  switch req.Method {
  case "POST":
    var grade utils.Grade
    err := json.NewDecoder(req.Body).Decode(&grade)
    if err != nil {
      http.Error(res, err.Error(), http.StatusInternalServerError)
      return
    }

    log.Println(grade, grade.NameStudent, grade.Subject, grade.Grade)
    err = students.AddGrade(&grade)
    if err != nil {
      http.Error(res, err.Error(), http.StatusInternalServerError)
      return
    }

    setJSONResponse(res, []byte("Ok"))

  default:
    err := errors.New(UnhandledRequest)
    http.Error(res, err.Error(), http.StatusInternalServerError)
  }
}

func getStudents(res http.ResponseWriter, req *http.Request) {
  switch req.Method {
  case "GET":
    json, err := json.MarshalIndent(students, "", "  ")
    if err != nil {
      log.Println(err.Error())
      http.Error(res, err.Error(), http.StatusInternalServerError)
      return
    }

    setJSONResponse(res, json)

  default:
    err := errors.New(UnhandledRequest)
    http.Error(res, err.Error(), http.StatusInternalServerError)
  }
}

func getStudent(res http.ResponseWriter, req *http.Request) {

  switch req.Method {
  case "GET":
    id, err := strconv.ParseUint(
      strings.TrimPrefix(req.URL.Path, "/GET/"), 10, 64)

    if err != nil {
      http.Error(res, err.Error(), http.StatusInternalServerError)
      return
    }

    student, err := students.GetStudent(int64(id))
    if err != nil {
      http.Error(res, err.Error(), http.StatusInternalServerError)
      return
    }

    log.Println(student)
    json, err := json.MarshalIndent(student, "", "  ")

    if err != nil {
      http.Error(res, err.Error(), http.StatusInternalServerError)
      return
    }

    setJSONResponse(res, json)

    default:
      err := errors.New(UnhandledRequest)
      http.Error(res, err.Error(), http.StatusInternalServerError)
  }
}

/*
func studentsIdHandler(res http.ResponseWriter, req *http.Request) {
  // TODO
  log.Println(req)
}
*/

func main() {
  students = utils.NewDataStudents()
  http.HandleFunc("/POST_calificacion", addGrade)
  http.HandleFunc("/GET_estudiantes", getStudents)
  http.HandleFunc("/GET/", getStudent)
  // http.HandleFunc("/estudiantes", studentsIdHandler)
  log.Println("Starting the server...")
  http.ListenAndServe(":" + labs.PORT, nil)
}

func setJSONResponse(res http.ResponseWriter, res_json []byte) {
  res.Header().Set(
    "Content-Type",
    "application/json",
  )

  res.Write(res_json)
}
