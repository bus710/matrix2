package main

// webSocketMessage - can be used for the websocket response to the clients
type webSocketMessage struct {
	Type string `json:"type"`
	Data string `json:"data"`
}

// Matrix data
type matrixData struct {
	point []pointData
}

type pointData struct {
	pointColor []int
}
