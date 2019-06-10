/* About sensorHat.go
This module has been developed to access the HAT (of course...). */

package main

import (
	"log"
	"runtime"
	"strings"
	"sync"
	"time"

	"periph.io/x/periph/conn"
	"periph.io/x/periph/conn/i2c"
	"periph.io/x/periph/conn/i2c/i2creg"
	"periph.io/x/periph/host"
)

// The main struc of this module
type senseHat struct {
	// app-wide items
	wait          *sync.WaitGroup
	chanStop      chan bool
	chanDataReady chan bool

	// local items
	isARM  bool
	i2cBus i2c.BusCloser
	i2cDev i2c.Dev
	i2cCon conn.Conn

	// buffers for the matrix
	matrixAddr uint16
	dotIndex   uint8
	bufRaw     [193]byte
	bufR       [64]byte
	bufG       [64]byte
	bufB       [64]byte
}

// init - assigns data and channels
func (sh *senseHat) init(wait *sync.WaitGroup) {
	log.Println()
	sh.chanStop = make(chan bool, 1)
	sh.chanDataReady = make(chan bool, 1)
	sh.wait = wait

	// To make sure the arch afterwords
	if strings.Contains(runtime.GOARCH, "arm") {
		// To check the archteture afterwords
		sh.isARM = true

		// To initialize the baseline drivers
		_, err := host.Init()
		if err != nil {
			log.Println(err)
		}

		// To open the i2c of RPI
		bus, err := i2creg.Open("")
		if err != nil {
			log.Println(err)
		}

		// To initialize some numbers
		sh.matrixAddr = uint16(0x0046) // SensorHat's AVR MCU uses 0x46 for the matrix
		sh.dotIndex = 0

		// To initialize the i2c bus
		sh.i2cBus = bus
		// To avoid Vet's warning, the specific keys are being used here
		sh.i2cDev = i2c.Dev{Bus: sh.i2cBus, Addr: sh.matrixAddr}
		sh.i2cCon = &sh.i2cDev

		err = sh.display()
		if err != nil {
			log.Println("Cannot use the i2c bus")
		}

	} else {
		// If the arch is not ARM...
		sh.isARM = false
	}
}

// run - runs the main go routine
func (sh *senseHat) run() {
	tick := time.Tick(1000 * time.Millisecond)

	if sh.isARM {
		defer sh.i2cBus.Close()
	}

StopFlag:
	for {
		select {
		case <-sh.chanStop:
			// To shutdown gracefully.
			// Some cleaning action can be added here.
			log.Println("got a signal from the chanStop")
			break StopFlag
		case <-sh.chanDataReady:
			// When the webserver safely received a chunk of data
			if sh.isARM {
				log.Println("data ready")
				err := sh.display()
				if err != nil {
					log.Println("error is occurred")
				}
			}
		case <-tick:
			// To run some task periodically
			// log.Println("test from the sensorhat routine")
		}
	}
	sh.wait.Done()
}

func (sh *senseHat) display() (err error) {

	// To set a certain pixel
	// sh.dotIndex++
	// if sh.matrixAddr > 63 {
	// 	sh.dotIndex = 0
	// }
	// sh.bufR[sh.dotIndex] = 0x20
	// sh.bufG[sh.dotIndex] = 0x00
	// sh.bufB[sh.dotIndex] = 0x00
	// sh.bufRaw[sh.dotIndex] = 0x00

	// Actual mapping
	j := int(0)
	for i := 0; i < 64; i++ {
		j = int(i/8) * 8
		j = j + j
		sh.bufRaw[i+j+1] = sh.bufR[i] / 4
		sh.bufRaw[i+j+9] = sh.bufG[i] / 4
		sh.bufRaw[i+j+17] = sh.bufB[i] / 4
	}

	// Actual writing
	writtenData, err := sh.i2cDev.Write(sh.bufRaw[:])
	if err != nil {
		return err
	} else if writtenData != 193 {
		return err
	}

	log.Println(writtenData, "bytes were written to the matrix")
	// log.Println(sh.bufRaw)

	// sh.bufR[sh.dotIndex] = 0x00
	// sh.bufG[sh.dotIndex] = 0x00
	// sh.bufB[sh.dotIndex] = 0x00

	return nil
}
