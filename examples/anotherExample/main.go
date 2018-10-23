package main

import (
	"log"
	"time"

	l "simonTest/logger"
)

var logger = l.New("anotherExample", "Main", true)

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
			logger.Debugf("This is a DEBUG message", "And here's some additiaonal information")
		}
	}()
}
