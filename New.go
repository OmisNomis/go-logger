package logger

import (
	"encoding/json"
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
		sock := fmt.Sprintf(
			"/tmp/sockets/%s.%s.sock",
			strings.ToUpper(p), strings.ToUpper(s),
		)

		ln, err := net.Listen("unix", sock)
		if err != nil {
			log.Fatalf("LOGGER: listen error: %+v", err)
		}

		ch := make(chan os.Signal, 1)
		signal.Notify(ch, os.Interrupt, syscall.SIGTERM)

		logger.handleShutdown(ln, ch)

		for {
			fd, err := ln.Accept()
			if err != nil {
				log.Fatalf("LOGGER: Accept error: %+v", err)
			}

			logger.handleIncomingMessages(fd)
		}

	}()

	return &logger
}

func (l *Logger) handleShutdown(ln net.Listener, c chan os.Signal) {
	// Shut down the socket if the application closes
	go func() {
		<-c
		log.Printf("LOGGER: Shutting down unix socket for Logger")
		ln.Close()
		os.Exit(0)
	}()
}

func (l *Logger) handleIncomingMessages(c net.Conn) {
	go func() {
		for {
			buf := make([]byte, 512)
			nr, err := c.Read(buf)
			if err != nil {
				return
			}

			data := buf[0:nr]
			op := strings.Split(strings.ToLower(string(data)), "\n")[0]

			switch op {
			case "debug":
				l.Conf.mu.Lock()
				l.Conf.Debug = !l.Conf.Debug
				l.Conf.mu.Unlock()
			case "list":
				b, err := json.Marshal(l)
				_, err = c.Write(b)
				if err != nil {
					log.Fatal("Writing client error: ", err)
				}
			}
		}
	}()
}
