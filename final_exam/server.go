package main

import (
  "log"
  "net/http"
  "labs"
  "labs/utils"
  "fmt"
  "encoding/json"
  "errors"
  "strconv"
)

var chatA * utils.Chat

type Payload struct {
  Messages  []utils.Message `json:"Messages"`
}

func NewPayload(content []utils.Message) * Payload {
  return &Payload{content}
}

const (
  UnhandledRequest = "Unhandled request."
)

func formatJSONResponse(msg string) []byte {
  res := fmt.Sprintf(`{"code": "%s"}`, msg)
  return []byte(res)
}

func root(res http.ResponseWriter, req * http.Request) {
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

func getStyle(res http.ResponseWriter, req *http.Request) {
  res.Header().Set(
    "Content-Type",
    "text/css",
  )

  content, err := labs.ReadFileContent("style/style.css")
  if err != nil { log.Fatal("Read the file content") }

  fmt.Fprintf(
    res,
    content,
  )
}

func getScriptUtils(res http.ResponseWriter, req *http.Request) {
  res.Header().Set(
    "Content-Type",
    "text/javascript",
  )

  content, err := labs.ReadFileContent("js/utils.js")
  if err != nil { log.Fatal(err, "Read the file content") }

	fmt.Fprintf(
		res,
    content,
	)
}

func getScriptRequests(res http.ResponseWriter, req *http.Request) {
  res.Header().Set(
    "Content-Type",
    "text/javascript",
  )

  content, err := labs.ReadFileContent("js/api_requests.js")
  if err != nil { log.Fatal("Read the file content") }

	fmt.Fprintf(
		res,
    content,
	)
}

func handleMessages(res http.ResponseWriter, req *http.Request) {
  switch req.Method {
  case "POST":
    var msg utils.Message

    err := json.NewDecoder(req.Body).Decode(&msg)
    if err != nil {
      http.Error(res, err.Error(), http.StatusInternalServerError)
      return
    }

    chatA.Msgs = append(chatA.Msgs, msg)
    jsonMsg := formatJSONResponse("Ok")

    data := chatA.GetMessages()
    log.Println(data)
    setJSONResponse(res, jsonMsg)

  case "GET":
    if err := req.ParseForm(); err != nil {
			log.Println(err)
			return
		}

    // Get the index from the parameter and parse it to int64
    strParam := req.FormValue("lastIndex")
    lastIndex, err := strconv.ParseInt(strParam, 10, 64)
    if err != nil {
      http.Error(res, err.Error(), http.StatusInternalServerError)
      return
    }

    response := *NewPayload(chatA.GetFromIndex(lastIndex))

    json, err := json.MarshalIndent(response, "", "  ")
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

func getChat(res http.ResponseWriter, req * http.Request) {
  switch req.Method {
	case "POST":
		if err := req.ParseForm(); err != nil {
			log.Println(err)
			return
		}

    // Get the chat name and the username
    chatName := req.FormValue("chat")
    username := req.FormValue("username")

    htmlContent, err := labs.ReadFileContent("chat.html")
    if err != nil {
      log.Println(err)
      return
    }

    res.Header().Set(
      "Content-Type",
      "text/html",
    )

    fmt.Fprintf(
      res,
      htmlContent,
      chatName,
      username,
    )
	}
}

func main() {
  chatA = new(utils.Chat)

  handler := http.NewServeMux()
  handler.HandleFunc("/", root)

  handler.HandleFunc("/style/style.css", getStyle)
  handler.HandleFunc("/js/api_requests.js", getScriptRequests)
  handler.HandleFunc("/js/utils.js", getScriptUtils)

  handler.HandleFunc("/chat", getChat)
  handler.HandleFunc("/API/messages", handleMessages)


  log.Println("Starting the server")
  address := ":" + labs.PORT
  http.ListenAndServe(address, handler)
}

func setJSONResponse(res http.ResponseWriter, res_json []byte) {
  res.Header().Set(
    "Content-Type",
    "application/json",
  )

  res.Write(res_json)
}
