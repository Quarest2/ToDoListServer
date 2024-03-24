package host

import (
	"ToDoList/ToDoListServer/TestClient/Tasks"
	"encoding/json"
	"fmt"
	log "github.com/sirupsen/logrus"
	"net"
	"time"
)

func (h *Host) Connect() {
	var err error

	h.mu.Lock()

	log.Printf("Connecting to %s:%d", h.Domain, h.Port)

	if h.Conn != nil {
		h.Conn.Close()
	}

	h.Conn = nil

	for h.Conn == nil {
		h.Conn, err = net.Dial("tcp", fmt.Sprintf("%s:%d", h.Domain, h.Port))

		if err != nil {
			h.Conn = nil

			log.Printf("Failed to connect to %s:%d", h.Domain, h.Port)
			time.Sleep(1 * time.Second)
			continue
		}

		log.Printf("Connected to %s:%d", h.Domain, h.Port)

		go h.HandleConnection()
		h.mu.Unlock()
	}
}

func (h *Host) HandleConnection() {
	buf := make([]byte, 1024)

	for {
		_, err := h.Conn.Read(buf)
		if err != nil {
			log.Printf("Failed to read from %s", h.Domain)
			h.Connect()
			return
		}

		log.Printf("Received package from %s", h.Domain)
		var task Tasks.Task
		err = json.Unmarshal(buf, &task)

		if err != nil {
			log.Printf("Failed to unmarshal JSON from %s", h.Domain)
			continue
		}
	}
}

func (h *Host) Broadcast(message []byte) {
	if !h.mu.TryLock() {
		return
	}

	_, err := h.Conn.Write(message)

	h.mu.Unlock()

	if err != nil {
		log.Printf("Failed to send message to %s", h.Domain)
		log.Printf("[ERR]: %s", err.Error())
		h.Connect()
		log.Printf("Reconnected to %s", h.Domain)
	}
}

func (h *Host) Close() {
	log.Printf("Closing connection to %s", h.Domain)

	h.mu.Lock()

	if h.Conn != nil {
		h.Conn.Close()
	}

	h.mu.Unlock()
}
