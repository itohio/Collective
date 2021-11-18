package rpc

//go:generate protoc -I=../../../schema/protobuf -I=${GOPATH}/pkg/mod/ -I=${GOPATH}/src --gogofaster_out=. schema.proto

import (
	"log"
	"net/http"
	"time"

	"github.com/itohio/collective/backend/pkg/db"

	"github.com/go-chi/chi/v5"
	"github.com/gorilla/websocket"
)

const (
	writeWait      = 10 * time.Second
	pongWait       = 60 * time.Second
	pingPeriod     = (pongWait * 9) / 10
	maxMessageSize = 1024 * 1024
)

type rpcHandler struct {
	upgrader websocket.Upgrader
}

func New(db db.DBType) http.Handler {
	router := chi.NewRouter()

	router.Handle("/{session}", &rpcHandler{
		upgrader: websocket.Upgrader{
			ReadBufferSize:  1024,
			WriteBufferSize: 1024,
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
		},
	})

	return router
}

type client struct {
	conn    *websocket.Conn
	session string

	// Buffered channel of outbound messages.
	send chan []byte
}

// Handler handles websocket requests from the peer.
func (h *rpcHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	conn, err := h.upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	session := chi.URLParam(r, "session")

	// TODO: Validate session

	c := &client{
		send:    make(chan []byte, 256),
		conn:    conn,
		session: session,
	}

	go c.writePump()
	go c.readPump()
}

// readPump pumps messages from the websocket connection to the hub.
func (c client) readPump() {
	defer func() {
		c.conn.Close()
	}()

	c.conn.SetReadLimit(maxMessageSize)
	c.conn.SetReadDeadline(time.Now().Add(pongWait))
	c.conn.SetPongHandler(func(string) error { c.conn.SetReadDeadline(time.Now().Add(pongWait)); return nil })

	for {
		_, msg, err := c.conn.ReadMessage()

		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway) {
				log.Printf("error: %v", err)
			}
			break
		}

		// TODO pass it to nats
		_ = msg
	}
}

// writePump pumps messages from nats to the websocket connection.
func (c *client) writePump() {
	ticker := time.NewTicker(pingPeriod)

	defer func() {
		ticker.Stop()
		c.conn.Close()
	}()

	for {
		select {
		case message, ok := <-c.send:
			if !ok {
				c.write(websocket.CloseMessage, []byte{})
				return
			}

			if err := c.write(websocket.TextMessage, message); err != nil {
				return
			}
		case <-ticker.C:
			if err := c.write(websocket.PingMessage, []byte{}); err != nil {
				return
			}
		}
	}
}

// write writes a message with the given message type and payload.
func (c *client) write(mt int, payload []byte) error {
	c.conn.SetWriteDeadline(time.Now().Add(writeWait))
	return c.conn.WriteMessage(mt, payload)
}
