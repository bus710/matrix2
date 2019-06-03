package main

import (
	"log"
	"sync"
)

func main() {
	log.Println("Hello!")

	waitInstance := sync.WaitGroup{}
	serverInstance := webServer{}
	signalInstance := termSignal{}

	signalInstance.init(&waitInstance, &serverInstance)
	serverInstance.init(&waitInstance)

	waitInstance.Add(1)
	go signalInstance.catcher()
	waitInstance.Add(1)
	go serverInstance.run()

	waitInstance.Wait()

	log.Println()
	log.Println("See you again!")
}
