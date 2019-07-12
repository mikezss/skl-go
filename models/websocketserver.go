package models

import (
	_ "bytes"

	"encoding/json"
	_ "errors"
	_ "strconv"
	"strings"

	_ "github.com/gorilla/websocket"
)

const (
	OpenWebsocket   = 1 //
	CloseWebsocket  = 2 //
	ReceivedMessage = 3 //
	SendedMessage   = 4 //
)

var ws *WebsocketServer

type SynDataMessage struct {
	Messagetype int    `json:"Messagetype"`
	Ip          string `json:"Ip"`
	Message     DATAMOVEAPPLYITEM
}
type WebsocketServer struct {
	// Registered clients.
	Clients map[*WebsocketClient]bool

	// Inbound messages from the clients.
	Broadcast chan []byte

	// Register requests from the clients.
	Register chan *WebsocketClient

	// Unregister requests from clients.
	Unregister chan *WebsocketClient

	//unread data message
	Unreaddatamessage []*SynDataMessage

	//指定IP data message
	Datamessage chan *SynDataMessage
}

func NewWebsocketServer() *WebsocketServer {
	Getlog().Debug("NewWebsocketServer() start==>")
	if ws == nil {
		ws = &WebsocketServer{
			Clients:           make(map[*WebsocketClient]bool),
			Broadcast:         make(chan []byte),
			Register:          make(chan *WebsocketClient),
			Unregister:        make(chan *WebsocketClient),
			Unreaddatamessage: make([]*SynDataMessage, 0),
			Datamessage:       make(chan *SynDataMessage),
		}
		return ws
	} else {
		return ws
	}

}
func GetWebsocketServer() *WebsocketServer {
	Getlog().Debug("GetWebsocketServer() start==>")
	if ws == nil {
		return NewWebsocketServer()
	} else {
		return ws
	}
}
func (ws *WebsocketServer) Run() {
	Getlog().Debug("Run() start==>")
	for {
		select {
		case client := <-ws.Register:
			ws.Clients[client] = true
		case client := <-ws.Unregister:
			if _, ok := ws.Clients[client]; ok {
				delete(ws.Clients, client)
				close(client.Send)
			}
		case message := <-ws.Broadcast:
			for client, _ := range ws.Clients {
				select {
				case client.Send <- message:
				default:
					close(client.Send)
					delete(ws.Clients, client)
				}
			}

		case datamessage := <-ws.Datamessage:
			isclientconn := false
			for client := range ws.Clients {
				ip := client.Conn.RemoteAddr().String()
				//Getlog().Debug("ip==>" + ip)
				//Getlog().Debug("datamessage.Ip==>" + datamessage.Ip)
				iparr := strings.Split(ip, ":")
				dataiparr := strings.Split(datamessage.Ip, ":")
				//Getlog().Debug("iparr0==>" + iparr[0])
				//Getlog().Debug("dataiparr0==>" + dataiparr[0])
				if iparr[0] == dataiparr[0] {
					isclientconn = true
					databyte, err := json.Marshal(datamessage)
					if err != nil {
						client.Send <- []byte(err.Error())
					}
					client.Send <- databyte
					break
				}

			}
			if !isclientconn {
				ws.Unreaddatamessage = append(ws.Unreaddatamessage, datamessage)
			}
		}

	}
}
func (ws *WebsocketServer) PushUnreaddatamessage(wc *WebsocketClient) {
	Getlog().Debug("PushUnreaddatamessage start==>")
	for idx, datamessage := range ws.Unreaddatamessage {
		databyte, err := json.Marshal(datamessage)
		if err != nil {
			Getlog().Error("json.Marshal(datamessage)==>" + err.Error())
			continue
		}
		wc.Send <- databyte
		if idx < len(ws.Unreaddatamessage)-1 {
			ws.Unreaddatamessage = append(ws.Unreaddatamessage[:idx], ws.Unreaddatamessage[idx+1:]...)
		} else {
			ws.Unreaddatamessage = ws.Unreaddatamessage[:idx]
		}

	}
}
