package main

import (
  "log"
  "net/http"
  "labs"
  "fmt"
)

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

func getChat(res http.ResponseWriter, req * http.Request) {
  switch req.Method {
	case "POST":
		if err := req.ParseForm(); err != nil {
			fmt.Fprintf(res, "ParseForm() error %v", err)
			return
		}

		log.Println(req.PostForm)

    // Get the chat name and the username
    chatName := req.FormValue("chat")
    username := req.FormValue("username")

    log.Println(username, chatName)
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
      string(chatName + " " + username),
    )
	}
}

func main() {
  handler := http.NewServeMux()
  handler.HandleFunc("/", root)
  handler.HandleFunc("/chat", getChat)

  log.Println("Starting the server")
  address := ":" + labs.PORT
  http.ListenAndServe(address, handler)
}
