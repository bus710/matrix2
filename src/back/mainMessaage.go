package main

// Message - can be used for the websocket response to the clients
type webSocketMessage struct {
	Type string `json:"type"`
	Data string `json:"data"`
}
