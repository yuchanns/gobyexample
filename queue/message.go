package main

type Message struct {
	Content []string `json:"message"`
	Id      int64    `json:"id"`
}
