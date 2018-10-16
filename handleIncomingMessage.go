package logger

import (
	"fmt"
	"log"
	"net"
	"strings"
)

func (l *Logger) handleIncomingMessage(c net.Conn) {
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
				l.handleDebugChange()

				res := fmt.Sprintf("Debug set to %t\n", l.Conf.Debug)
				_, err = c.Write([]byte(res))
				if err != nil {
					log.Printf("Writing client error: %+v", err)
				}
			case "list":
				res := fmt.Sprintf("Debug currently set to %t\n", l.Conf.Debug)
				_, err = c.Write([]byte(res))
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
