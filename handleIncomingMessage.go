package logger

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func (l *Logger) handleIncomingMessage(c net.Conn) {
	go func() {
		scanner := bufio.NewScanner(c)
		for scanner.Scan() {
			cmd := scanner.Text()
			switch cmd {
			case "debug":
				// TODO - add REGEX support
				l.handleDebugChange()
				l.sendDebugStatus(c)
			case "trace":
				// TODO - Add in an option to send debug messages to socket
				log.Println("Send debug messages to the socket and not a log file")
			case "status":
				l.sendDebugStatus(c)
			case "quit":
				c.Close()
				return
			case "help":
				_, err := c.Write([]byte("Available Commands:\n\t\t-- debug\n\t\t-- status\n\t\t-- help (this command)\n"))
				if err != nil {
					log.Printf("Writing client error: %+v", err)
				}
			default:
				_, err := c.Write([]byte(fmt.Sprintf("Command not found: %s\n", cmd)))
				if err != nil {
					log.Printf("Writing client error: %+v", err)
				}
			}
		}
	}()
}

func (l *Logger) handleDebugChange() {
	l.Conf.mu.Lock()
	defer l.Conf.mu.Unlock()

	l.Conf.Debug = !l.Conf.Debug
}

func (l *Logger) sendDebugStatus(c net.Conn) {
	l.Conf.mu.RLock()
	defer l.Conf.mu.RUnlock()

	res := fmt.Sprintf("Debug set to %t\n", l.Conf.Debug)
	_, err := c.Write([]byte(res))
	if err != nil {
		log.Printf("Writing client error: %+v", err)
	}
}
