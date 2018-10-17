package logger

import (
	"fmt"
	"log"
	"math/rand"
	"net"
	"os"
	"os/signal"
	"strings"
	"sync"
	"syscall"
	"time"
)

func init() {
	err := os.MkdirAll("/tmp/sockets", os.ModePerm)
	if err != nil {
		log.Printf("ERROR: Unable to make stats directory %s: %+v", "/tmp/sockets", err)
	}
}

// New comment
func New(packageName string, serviceName string, enableColours bool) *Logger {
	logger := Logger{
		PackageName: strings.ToUpper(packageName),
		ServiceName: strings.ToUpper(serviceName),
		Colour:      enableColours,
		Conf: Config{
			mu: new(sync.RWMutex),
			Trace: TraceSettings{
				Sockets: make(map[net.Conn]string),
			},
		},
	}

	// Run a listener on a Unix socket
	go func() {
		n := fmt.Sprintf(
			"/tmp/sockets/%s.%s%d.sock",
			strings.ToUpper(packageName), strings.ToUpper(serviceName), getRand(),
		)

		ln, err := net.Listen("unix", n)
		if err != nil {
			log.Fatalf("LOGGER: listen error: %+v", err)
		}

		log.Printf("Socket created. Connect with 'nc -U %s'", n)

		ch := make(chan os.Signal, 1)
		signal.Notify(ch, os.Interrupt, syscall.SIGTERM)

		logger.handleShutdown(ln, ch)

		for {
			fd, err := ln.Accept()
			if err != nil {
				log.Fatalf("LOGGER: Accept error: %+v", err)
			}

			logger.handleIncomingMessage(fd)
		}

	}()

	return &logger
}

func getRand() int {
	rand.Seed(time.Now().UnixNano())
	min := 100000000
	max := 999999999
	return rand.Intn(max-min) + min
}
