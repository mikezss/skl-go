package models

import (
	"bytes"

	"encoding/json"
	_ "errors"
	"time"

	"github.com/gorilla/websocket"
)

const (
	// Time allowed to write a message to the peer.
	writeWait = 10 * time.Second

	// Time allowed to read the next pong message from the peer.
	pongWait = 60 * time.Second

	// Send pings to peer with this period. Must be less than pongWait.
	pingPeriod = (pongWait * 9) / 10

	// Maximum message size allowed from peer.
	maxMessageSize = 512
)

var (
	newline = []byte{'\n'}
	space   = []byte{' '}
)

type WebsocketClient struct {
	Ws *WebsocketServer

	// The websocket connection.
	Conn *websocket.Conn

	// Buffered channel of outbound messages.
	Send chan []byte

	// Buffered channel of inbound messages.
	Recieve chan []byte
}

func (c *WebsocketClient) ReadMessage() {
	Getlog().Debug("ReadMessage() start==>")
	//c.Conn.SetReadLimit(maxMessageSize)
	//c.Conn.SetReadDeadline(time.Now().Add(pongWait))
	//c.Conn.SetPongHandler(func(string) error { c.conn.SetReadDeadline(time.Now().Add(pongWait)); return nil })
	for {
		messagetype, message, err := c.Conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				Getlog().Error("c.Conn.ReadMessage()==>" + err.Error())
			}
			break
		}
		if messagetype == websocket.CloseMessage {

			c.Ws.Unregister <- c
			c.Conn.Close()
			break

		}

		message = bytes.TrimSpace(bytes.Replace(message, newline, space, -1))
		c.Ws.Broadcast <- message
	}
}
func (c *WebsocketClient) WriteMessage() {
	Getlog().Debug("WriteMessage() start==>")
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		c.Conn.Close()
	}()
	for {
		select {
		case message, ok := <-c.Send:
			c.Conn.SetWriteDeadline(time.Now().Add(writeWait))
			if !ok {
				// The hub closed the channel.
				c.Conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			w, err := c.Conn.NextWriter(websocket.TextMessage)
			if err != nil {
				Getlog().Error("c.Conn.NextWriter(websocket.TextMessage)==>" + err.Error())
				return
			}
			_, err = w.Write(message)
			if err != nil {
				Getlog().Error("w.Write(message)==>" + err.Error())
				sdm := SynDataMessage{}
				err = json.Unmarshal(message, sdm)
				c.Ws.Unreaddatamessage = append(c.Ws.Unreaddatamessage, &sdm)

			}

			// Add queued chat messages to the current websocket message.
			n := len(c.Send)
			for i := 0; i < n; i++ {
				_, err := w.Write(<-c.Send)
				if err != nil {
					Getlog().Error("w.Write(<-c.Send)==>" + err.Error())
					sdm := SynDataMessage{}
					err = json.Unmarshal(message, sdm)
					c.Ws.Unreaddatamessage = append(c.Ws.Unreaddatamessage, &sdm)
				}
			}

			if err := w.Close(); err != nil {
				Getlog().Error("w.Close==>" + err.Error())
				return
			}
		case <-ticker.C:
			c.Conn.SetWriteDeadline(time.Now().Add(writeWait))
			if err := c.Conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				Getlog().Error("c.Conn.WriteMessage==>" + err.Error())
				return
			}
		}
	}
}
