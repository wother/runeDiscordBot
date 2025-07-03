package models

type Message struct {
    Content string `json:"content"`
    Author  string `json:"author"`
}

type User struct {
    ID       string `json:"id"`
    Username string `json:"username"`
    Avatar   string `json:"avatar"`
}