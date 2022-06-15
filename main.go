package main

import (
	"log"
	"net/http"

	"golang.org/x/net/websocket"
)

// Init function to register http handlers
func init() {
	http.Handle("/ws", websocket.Handler(HandleNewConnection))
	http.Handle("/", http.FileServer(http.Dir("./clientApp/public")))
}

func main() {
	log.Fatal(http.ListenAndServe(":3000", nil))
}
