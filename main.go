package main

import (
  "net/http"
  "github.com/laurent22/toml-go"
  "github.com/trevex/golem"
)

func main() {
  // Parse config file
  var parser toml.Parser
  doc := parser.ParseFile("config/app.conf")

  // WebSocket handler
  router := golem.NewRouter()
  initializeChat(router)

  // Web
  http.Handle("/", http.FileServer(http.Dir("./public")))
  http.HandleFunc("/ws", router.Handler())

  port := doc.GetString("web.port")
  err := http.ListenAndServe(":" + port, nil);
  if err != nil {
    panic(err)
  }
}
