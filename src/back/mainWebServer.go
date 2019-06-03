package main

import (
	"log"
	"net/http"
	"sync"
	"time"

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
func (s *webServer) init(
	wait *sync.WaitGroup) (err error) {

	s.wait = wait
	s.instance = &http.Server{Addr: ":3000"}
	return nil
}

// run - delivers the static web files and serves the REST API (+ websocket)
func (s *webServer) run() (err error) {
	// WebSocket
	http.Handle("/message", websocket.Handler(s.socket))

	// Web Contents
	http.Handle("/", http.FileServer(http.Dir("../front/build")))

	// Server up and running
	log.Println(s.instance.ListenAndServe())

	s.wait.Done()
	return nil
}

// socket - websocket handler
func (s *webServer) socket(ws *websocket.Conn) {
	log.Println(ws.Request().RemoteAddr)

	defer ws.Close()

	for {
		s.receivedItemWS = &webSocketMessage{}
		// receive a message using the codec
		if err := websocket.JSON.Receive(ws, &s.receivedItemWS); err != nil {
			log.Println(err)
			break
		}

		tmp := s.receivedItemWS.Message
		log.Println("Received message:", tmp)

		if tmp == "true" {
			s.responseItemWS = &webSocketMessage{"false"}
		} else {
			s.responseItemWS = &webSocketMessage{"true"}
		}

		// send a response
		if err := websocket.JSON.Send(ws, s.responseItemWS); err != nil {
			log.Println(err)
			break
		}

		log.Println("1 - ")

		time.Sleep(time.Millisecond * 3000)

		s.responseItemWS = &webSocketMessage{"true"}

		// send a response
		if err := websocket.JSON.Send(ws, s.responseItemWS); err != nil {
			log.Println(err)
			break
		}

		log.Println("2 - ")

	}
}
