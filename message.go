package main

import "encoding/json"

var messages []*Message

// Mssage struct represents a message (Members are capitalized for use in JSON)
type Message struct {
	SenderID string `json:"senderId"`
	Username string `json:"username"`
	Message  string `json:"message"`
}

// Post method broadcast message to all clients and add to message slice
func (m *Message) Post() {
	m.Broadcast()
	messages = append(messages, m)
}

// Broadcast method ends message to all clients in client slice
func (m *Message) Broadcast() {
	for _, client := range clients {
		m.BroadcastTo(client)
	}
}

// BroadcastTo method sends message to passed client
func (m *Message) BroadcastTo(to *Client) {
	byteMessage, _ := json.Marshal(m)
	to.ws.Write(byteMessage)
}
