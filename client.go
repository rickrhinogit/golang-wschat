package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"

	"golang.org/x/net/websocket"
)

var clients []*Client

type Client struct {
	id       string
	ip       string
	username string
	ws       *websocket.Conn
}

// StartListening method listens for messages from client
func (client *Client) StartListening() {
	buffer := make([]byte, 512)

	for {
		n, err := client.ws.Read(buffer[0:])

		if err != nil {
			ReleaseConnection(client)

			// Create message
			exitMessage := Message{
				SenderID: "System",
				Username: "System",
				Message:  fmt.Sprintf("%s has left the chat", client.username),
			}

			// Send Message
			exitMessage.Post()
		} else {
			HandleInputMessage(client, buffer[0:n])
		}
	}
}

// HandleNewConnection function creates a new client and stores it
func HandleNewConnection(c *websocket.Conn) {
	log.Println("New Connection!")

	// Create a new Client object and add it to the clients array
	newClient := Client{
		id:       genUserId(),
		ip:       c.Request().RemoteAddr,
		username: "",
		ws:       c,
	}

	clients = append(clients, &newClient)

	newClient.StartListening()

}

// ReleaseConnection function: remove client from client slice, ensure connection closed
func ReleaseConnection(client *Client) {
	log.Println("Released Connection", client.username, client.id)

	// Remove from client slice
	index := -1
	for idx, val := range clients {
		if client == val {
			index = idx
			break
		}
	}

	if index >= 0 {
		clients = append(clients[:index], clients[index+1:]...)
	}

	log.Println("Connection count: ", len(clients))
	client.ws.Close()
}

// genUserId function generates a random 10 character Id
func genUserId() string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFHIJKLMNOPQRSTUVWXYZ")
	rand.Seed(time.Now().UnixNano())

	b := make([]rune, 10)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}

	return string(b)
}
