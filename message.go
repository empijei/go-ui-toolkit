package goUIToolKit

import (
	"encoding/json"
	"io"

	"golang.org/x/net/websocket"
)

type Message struct {
	Command  string
	Selector string
	Args     []string
	Payload  []byte
}

func (r *Runtime) Send(m *Message) {
	//TODO cache content if not yet connected
	err := websocket.JSON.Send(r.ws, m)
	if err != nil {
		//TODO
		panic(err)
	}
}

func runtimeSender(ws *websocket.Conn) {
	//TODO deliver cached content and updates

}
func runtimeReceiver(ws *websocket.Conn) {
	dec := json.NewDecoder(ws)
	var msg Message
	var err error
	for err == nil {
		err = dec.Decode(&msg)
		if err == io.EOF {
			//UI was closed.
			//TODO
		} else if err != nil {
			panic(err)
		} else {
			parseMessage(&msg)
		}
	}
}

func parseMessage(m *Message) {
	switch m.Command {
	case "getvalue":

	case "event":
	}
}
