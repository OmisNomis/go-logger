package logger

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"strings"
	"sync"
	"syscall"
)

func init() {
	err := os.MkdirAll("/tmp/sockets", os.ModePerm)
	if err != nil {
		log.Printf("ERROR: Unable to make stats directory %s: %+v", "/tmp/sockets", err)
	}
}

// New comment
func New(p string, s string) *Logger {
	logger := Logger{
		PackageName: strings.ToUpper(p),
		ServiceName: strings.ToUpper(s),
		Conf: Config{
			mu: new(sync.Mutex),
		},
	}

	// Run a listener on a Unix socket
	go func() {
		n := fmt.Sprintf(
			"/tmp/sockets/%s.%s.sock",
			strings.ToUpper(p), strings.ToUpper(s),
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
