package main

import (
	"log"
	"net/http"
	"strconv"
	"sync"

	"golang.org/x/net/websocket"
)

// webServer - the main struct of this module
type webServer struct {
	// app-wide items
	wait *sync.WaitGroup

	// web items
	instance *http.Server

	// data
	receivedItemWS *webSocketMessage
	responseItemWS *webSocketMessage
}

// init - initializes the data and structs
func (wserver *webServer) init(
	wait *sync.WaitGroup) (err error) {

	wserver.wait = wait
	wserver.instance = &http.Server{Addr: ":3000"}
	return nil
}

// run - delivers the static web files and serves the REST API (+ websocket)
func (wserver *webServer) run() (err error) {
	// WebSocket
	http.Handle("/message", websocket.Handler(wserver.socket))

	// Web Contents
	http.Handle("/", http.FileServer(http.Dir("../front/build")))

	// Server up and running
	log.Println(wserver.instance.ListenAndServe())

	wserver.wait.Done()
	return nil
}

// socket - websocket handler
func (wserver *webServer) socket(wsocket *websocket.Conn) {

	log.Println(wsocket.Request().RemoteAddr)
	chanData := make(chan string, 1)
	chanResponse := make(chan bool, 1)

	defer wsocket.Close()

	// Processing routine
	go func() {
		for {
			select {
			case data := <-chanData:
				log.Println(data)
				chanResponse <- true
			}
		}
	}()

	// Sending routine
	go func() {
		for {
			select {
			case res := <-chanResponse:
				wserver.responseItemWS = &webSocketMessage{
					Type: "response", Data: strconv.FormatBool(res)}
				websocket.JSON.Send(wsocket, wserver.responseItemWS)
			}
		}
	}()

	// Receiving routine
	for {
		wserver.receivedItemWS = &webSocketMessage{}
		// receive a message using the codec
		if err := websocket.JSON.Receive(
			wsocket, &wserver.receivedItemWS); err != nil {
			// log.Println(err)
			break
		} else {
			tmp := wserver.receivedItemWS.Type
			log.Println("Received message:", tmp)
			chanData <- wserver.receivedItemWS.Data
		}

	}

	log.Println("Websocket closed")
}
