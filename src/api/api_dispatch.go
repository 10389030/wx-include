package api

import (
	"encoding/xml"
	"io/ioutil"
	"log"
	"net/http"
)

type MsgHandler func(http.ResponseWriter, *Message)

type MessageRouter struct {
	handlers map[MessageRoute]MsgHandler
}

func (router *MessageRouter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		CheckServer(w, r)
	case "POST":
		data, _ := ioutil.ReadAll(r.Body)

		msg := &Message{}
		err := xml.Unmarshal(data, msg)

		if err != nil {
			log.Panic("xml unmarshal post data error: %#v", err)
		}

		if handler, ok := router.handlers[msg.MessageRoute]; ok {
			log.Printf("Handle Found: %#v, %v", msg.MessageRoute, handler)
			handler(w, msg)
		} else {
			log.Printf("Handle Not Found: %#v", msg.MessageRoute)
			if r.ProtoAtLeast(1, 1) {
				w.Header().Set("Connection", "close")
			}
			w.WriteHeader(http.StatusNotFound)
		}

	default:
		log.Printf("unhandle request method: %s", r.Method)
	}
}

func (mux *MessageRouter) HandleMsg(msg *MessageRoute, handler func(http.ResponseWriter, *Message)) {
	mux.handlers[*msg] = MsgHandler(handler)
}

// creator
func NewMessageRouter() *MessageRouter {
	return &MessageRouter{
		handlers: map[MessageRoute]MsgHandler{},
	}
}
