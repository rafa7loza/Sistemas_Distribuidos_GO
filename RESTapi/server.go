package main

import (
  "log"
  "fmt"
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

func formatJSONResponse(msg string) []byte {
  res := fmt.Sprintf(`{"code": "%s"}`, msg)
  return []byte(res)
}

func addGrade(res http.ResponseWriter, req *http.Request) {
  switch req.Method {
  case "POST":
    var grade utils.Grade
    err := json.NewDecoder(req.Body).Decode(&grade)
    if err != nil {
      http.Error(res, err.Error(), http.StatusInternalServerError)
      return
    }

    err = students.AddGrade(&grade)
    if err != nil {
      http.Error(res, err.Error(), http.StatusInternalServerError)
      return
    }

    msg := formatJSONResponse("Ok")
    setJSONResponse(res, msg)

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

func handleStudent(res http.ResponseWriter, req *http.Request) {
  _id, err := strconv.ParseUint(
    strings.TrimPrefix(req.URL.Path, "/estudiante/"), 10, 64)

  id := int64(_id)
  if err != nil {
    http.Error(res, err.Error(), http.StatusInternalServerError)
    return
  }

  switch req.Method {
  case "GET":
    student, err := students.GetStudent(id)
    if err != nil {
      http.Error(res, err.Error(), http.StatusInternalServerError)
      return
    }

    json, err := json.MarshalIndent(student, "", "  ")

    if err != nil {
      http.Error(res, err.Error(), http.StatusInternalServerError)
      return
    }

    setJSONResponse(res, json)

  case "DELETE":
    err := students.RemoveStudent(id)
    var msg []byte
    if err != nil {
      msg = formatJSONResponse(err.Error())
    } else {
      msg = formatJSONResponse("Ok")
    }

    setJSONResponse(res, msg)

  case "PUT":
    var grade utils.Grade
    err := json.NewDecoder(req.Body).Decode(&grade)
    if err != nil {
      http.Error(res, err.Error(), http.StatusInternalServerError)
      return
    }

    err = students.UpdateSubjectGrade(id, &grade)
    var msg []byte
    if err != nil {
      msg = formatJSONResponse(err.Error())
    } else {
      msg = formatJSONResponse("Ok")
    }

    setJSONResponse(res, msg)


  default:
    err := errors.New(UnhandledRequest)
    http.Error(res, err.Error(), http.StatusInternalServerError)
  }
}

func main() {
  students = utils.NewDataStudents()
  http.HandleFunc("/calificacion", addGrade)
  http.HandleFunc("/estudiantes", getStudents)
  http.HandleFunc("/estudiante/", handleStudent)
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
