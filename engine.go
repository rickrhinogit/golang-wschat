package main

import (
	"encoding/json"
	"fmt"
)

// HandleInputMessage function is the primary handler for the websocket connection which listens for messages from the client
func HandleInputMessage(from *Client, data []byte) {
	// Parse JSON
	var input map[string]string
	json.Unmarshal(data, &input)

	// Based on the action, handle it
	switch input["action"] {
	case "post_message":
		newMessage := Message{
			SenderID: from.id,
			Username: from.username,
			Message:  input["message"],
		}

		newMessage.Post()

	case "initial_connection":
		// Set client username
		from.username = input["username"]

		// Create and send message with clients id
		newIDMessage := Message{
			SenderID: "System",
			Username: "System",
			Message:  fmt.Sprintf("YOURID: %v", from.id),
		}

		newIDMessage.BroadcastTo(from)

		// Send client chat history
		chatHistory, _ := json.Marshal(messages)
		from.ws.Write(chatHistory)

		// Announce to chat that new user has arrived
		newMessageToRoom := Message{
			SenderID: "System",
			Username: "System",
			Message:  fmt.Sprintf("%s joined the chat", from.username),
		}

		newMessageToRoom.Post()

	}
}
