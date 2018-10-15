package main

import (
	"log"
	"time"

	l "simonTest/logger"
)

var logger = l.New("SeperateExample", "Main")

func main() {
	runTest()

	waitCh := make(chan bool)
	<-waitCh

	log.Println(logger)
}

func runTest() {
	ticker := time.NewTicker(3 * time.Second)
	go func() {
		for range ticker.C {
			logger.Debug("This is a DEBUG message", "And here's some additiaonal information")
		}
	}()
}
