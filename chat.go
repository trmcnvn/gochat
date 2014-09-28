package main

import (
  "github.com/trevex/golem"
  "fmt"
)

type User struct {
  Name string `json:"name"`
  conn *golem.Connection
}

type Message struct {
  Name string `json:"from"`
  Message string `json:"msg"`
}

type Chat struct {
  users []*User
}

func initializeChat(r *golem.Router) {
  chat := Chat{}
  r.On("hello", chat.join)
  r.On("client", chat.message)
  r.OnClose(chat.leave)
}

func (chat *Chat) join(conn *golem.Connection, data *User) {
  // add user to list
  data.conn = conn
  chat.users = append(chat.users, data)

  // send list of users to all clients
  for _, v := range chat.users {
    v.conn.Emit("users", chat.users)
  }

  fmt.Println(data.Name, "connected.")
}

func (chat *Chat) leave(conn *golem.Connection) {
  // remove user for users list
  for i, v := range chat.users {
    if v.conn == conn {
      chat.users = append(chat.users[:i], chat.users[i+1:]...)
      fmt.Println(v.Name, "disconnected.")
      break
    }
  }

  // send list of users to all clients
  for _, v := range chat.users {
    v.conn.Emit("users", chat.users)
  }
}

func (chat *Chat) message(conn *golem.Connection, data *Message) {
  for _, v := range chat.users {
    if v.conn != conn {
      v.conn.Emit("server", data)
    }
  }
  fmt.Println(data.Name, "said", data.Message)
}
