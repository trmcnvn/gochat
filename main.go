package main

import (
  "net/http"
  "github.com/laurent22/toml-go"
)

func main() {
  // Parse config file
  var parser toml.Parser
  doc := parser.ParseFile("config/app.conf")

  // Web
  http.Handle("/", http.FileServer(http.Dir("./public")))

  err := http.ListenAndServe(doc.GetString("web.host"), nil);
  if err != nil {
    panic(err)
  }
}
