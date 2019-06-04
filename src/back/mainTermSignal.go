package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

// termSignal - the main struct of this module
type termSignal struct {
	// app-wide items
	wait     *sync.WaitGroup
	server   *webServer
	senseHat *senseHat

	// channels
	sigterm  chan os.Signal
	chanStop chan bool
}

// init - takes a WG, instance of service, and channels of go routines
// and keeps the assigned params in its struct to access later.
func (sig *termSignal) init(
	wait *sync.WaitGroup,
	serverInstance *webServer,
	senseHatInstance *senseHat) {

	// To assign instances to the pointers
	sig.wait = wait
	sig.server = serverInstance
	sig.senseHat = senseHatInstance

	//
	sig.chanStop = make(chan bool, 1)
	sig.sigterm = make(chan os.Signal, 1)
}

// catcher - a handler to catch the interrupts from keyboard (CTRL+C)
// and gracefully shuts down.
func (sig *termSignal) catcher() {

	// To connect the keyboard signal to the channel.
	signal.Notify(sig.sigterm, syscall.SIGINT, syscall.SIGTERM)

	// The routine waits here for the keyboard interrupt.
	select {
	case received := <-sig.sigterm:
		log.Println()
		log.Println("Received a CTRL+C", received)
		if err := sig.cleanup(); err != nil {
			log.Println(err)
		}
	case <-sig.chanStop:
		log.Println()
		log.Println("Received a signal")
		if err := sig.cleanup(); err != nil {
			log.Println(err)
		}
	}
}

// Running a graceful shutdown.
func (sig *termSignal) cleanup() (err error) {

	log.Println("Cleanup - started")

	// To send a signal to the sensorHat's channel
	sig.senseHat.chanStop <- true

	time.Sleep(time.Microsecond * 100)

	// To call the shutdown method of the webserver
	ctx, cancel := context.WithTimeout(
		context.Background(), time.Millisecond*300)

	defer cancel()

	if err := sig.server.instance.Shutdown(ctx); err != nil {
		return err
	}

	// To decrease the wait group
	sig.wait.Done()
	log.Println("Cleanup - done")

	return nil
}
