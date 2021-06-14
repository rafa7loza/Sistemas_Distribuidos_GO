package main

import (
  "log"
  "net/http"
  "labs"
  "labs/utils"
  "fmt"
)

var chatA * utils.Chat

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

func getChat(res http.ResponseWriter, req * http.Request) {
  switch req.Method {
	case "POST":
		if err := req.ParseForm(); err != nil {
			fmt.Fprintf(res, "ParseForm() error %v", err)
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

    msg := chatA.GetMessages()

    res.Header().Set(
      "Content-Type",
      "text/html",
    )

    fmt.Fprintf(
      res,
      htmlContent,
      string(chatName + " " + username),
      msg,
    )
	}
}

func main() {
  chatA = new(utils.Chat)

  handler := http.NewServeMux()
  handler.HandleFunc("/", root)
  handler.HandleFunc("/chat", getChat)

  handler.HandleFunc("/style/style.css", getStyle)

  log.Println("Starting the server")
  address := ":" + labs.PORT
  http.ListenAndServe(address, handler)
}
