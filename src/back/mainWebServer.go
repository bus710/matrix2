package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"sync"

	"golang.org/x/net/websocket"
)

// webServer - the main struct of this module
type webServer struct {
	// app-wide items
	wait     *sync.WaitGroup
	senseHat *senseHat

	// web items
	instance *http.Server

	// data
	receivedItemWS *webSocketMessage
	responseItemWS *webSocketMessage
}

// init - initializes the data and structs
func (wserver *webServer) init(
	wait *sync.WaitGroup,
	senseHatInstance *senseHat) (err error) {

	wserver.wait = wait
	wserver.instance = &http.Server{Addr: ":3000"}
	wserver.senseHat = senseHatInstance
	return nil
}

// run - delivers the static web files and serves the REST API (+ websocket)
func (wserver *webServer) run() (err error) {
	// WebSocket
	http.Handle("/message", websocket.Handler(wserver.socket))

	// Web Contents
	// The frontend side should be built by webdev build --output build
	// Otherwise, the location will be ../front/build
	http.Handle("/", http.FileServer(http.Dir("../front/build/web")))

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
				/* for future unmarshling
				https://mholt.github.io/json-to-go/ */
				var dataList matrixData
				// log.Println("Processing routine: " + data)

				if err := json.Unmarshal([]byte(data), &dataList); err != nil {
					log.Println(err)
					chanResponse <- false
				} else {
					if len(dataList) == 64 {
						// log.Println(dataList[0][0])

						for i := 0; i < 64; i++ {
							wserver.senseHat.bufR[i] = byte(dataList[i][0])
							wserver.senseHat.bufG[i] = byte(dataList[i][1])
							wserver.senseHat.bufB[i] = byte(dataList[i][2])
						}

						// To notify the data is ready to the sensorHat routine
						wserver.senseHat.chanDataReady <- true
						// To notify the data is ready to the client
						chanResponse <- true
					} else {
						chanResponse <- false
					}
				}
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

	wserver.receivedItemWS = &webSocketMessage{}
	// Receiving routine
	for {
		// receive a message using the codec
		if err := websocket.JSON.Receive(
			wsocket, &wserver.receivedItemWS); err != nil {
			// log.Println(err)
			break
		} else {
			messageType := wserver.receivedItemWS.Type
			messageData := wserver.receivedItemWS.Data
			log.Println("Received message type:", messageType)
			// log.Println("Received message data:", messageData)
			chanData <- messageData
		}
	}

	log.Println("Websocket closed")
}
