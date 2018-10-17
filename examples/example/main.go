package main

import (
	"errors"
	"log"
	"time"

	l "simonTest/logger"
)

var logger = l.New("Example", "Main", true)

func main() {
	runTest()

	waitCh := make(chan bool)
	<-waitCh

	log.Println(logger)
}

func runTest() {
	ticker1 := time.NewTicker(2 * time.Second)
	go func() {
		for range ticker1.C {
			logger.Debugf("Service A HAD A PROBLEM %s", "And here's some additiaonal information")
		}
	}()

	ticker2 := time.NewTicker(2 * time.Second)
	go func() {
		for range ticker2.C {
			logger.Debugf("Service 12522 HAD A PROBLEM %s", "And here's some additiaonal information")
		}
	}()

	ticker3 := time.NewTicker(2 * time.Second)
	go func() {
		for range ticker3.C {
			logger.Debugf("Service JimmyJammy HAD A PROBLEM %s", "And here's some additiaonal information")
		}
	}()

	ticker4 := time.NewTicker(2 * time.Second)
	go func() {
		for range ticker4.C {
			logger.Debugf("Service asgljsglsij HAD A PROBLEM %s", "And here's some additiaonal information")
		}
	}()

	ticker5 := time.NewTicker(2 * time.Second)
	go func() {
		for range ticker5.C {
			logger.Debugf("Service flookyprincess HAD A PROBLEM %s", "And here's some additiaonal information")
		}
	}()

	ticker6 := time.NewTicker(4 * time.Second)
	go func() {
		for range ticker6.C {
			logger.Errorf("Service flookyprincess HAD A PROBLEM %s", errors.New("OHH NOOO"))
		}
	}()

	ticker10 := time.NewTicker(4 * time.Second)
	go func() {
		for range ticker10.C {
			logger.Errorf("Service JimmyJammy HAD A PROBLEM")
		}
	}()
	ticker7 := time.NewTicker(5 * time.Second)
	go func() {
		for range ticker7.C {
			logger.Infof("Service flookyprincess Info")
		}
	}()
	ticker8 := time.NewTicker(5 * time.Second)
	go func() {
		for range ticker8.C {
			logger.Warnf("Service JimmyJammy Warning")
		}
	}()
}
