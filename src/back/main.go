package main

import (
	"log"
	"sync"
)

func main() {
	log.Println("Hello!")

	waitInstance := sync.WaitGroup{}
	senseHatInstance := senseHat{}
	serverInstance := webServer{}
	signalInstance := termSignal{}

	senseHatInstance.init(&waitInstance)
	serverInstance.init(&waitInstance, &senseHatInstance)
	signalInstance.init(&waitInstance, &serverInstance, &senseHatInstance)

	waitInstance.Add(1)
	go signalInstance.catcher()
	waitInstance.Add(1)
	go senseHatInstance.run()
	waitInstance.Add(1)
	go serverInstance.run()

	waitInstance.Wait()

	log.Println()
	log.Println("See you again!")
}
