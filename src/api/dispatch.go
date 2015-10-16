package api

import (
	"net/http"
	"log"
	"strings"
	"io/ioutil"
	"encoding/xml"
)

// Toolkit to transfer weixin Message to routing PATH
type PATHGenerator interface {
	ToPath() string
}

type MsgRouteInfo struct {
	MsgType 	string      `xml:"MsgType"`
	Event 		string 		`xml:"Event"`
	Data        []byte
}

func (msg *MsgRouteInfo) ToPath() (rst string) {
	segs := []string{
		"",
		strings.ToLower(msg.MsgType),
		strings.ToLower(msg.Event),
	}

	return strings.Join(segs, "/")
}

type MsgHandler func(http.ResponseWriter, *Message)
func (handler MsgHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
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
			
			log.Printf("POST: %#v", msg)	
			handler(w, msg)
		default:
			log.Printf("unhandle request method: %s", r.Method)
	}
}

// Toolkit to Handle PATHGenerator as URL &
// Handle URL ending with '/' as subtree rooting.
// ATTENTION: 
//   1 This will change the reuqest.URL.Path
//   2 Pattern must begin with '/'
type ServeMuxEx struct {
	*http.ServeMux
}

func (mux *ServeMuxEx) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.RequestURI == "*" {
		if r.ProtoAtLeast(1, 1) {
			w.Header().Set("Connection", "close")
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	h, pattern := mux.Handler(r)
	var patternLen = len(pattern)
	if patternLen > 0 && pattern[patternLen - 1] == '/' { 
		r.URL.Path = r.URL.Path[len(pattern) - 1:]
	}

	h.ServeHTTP(w, r)
}

func (mux *ServeMuxEx) HandleMsg(msg *MsgRouteInfo, handler func(http.ResponseWriter, *Message)) {
	mux.Handle(msg.ToPath(), MsgHandler(handler))
}

// creator
func NewServeMuxEx() (*ServeMuxEx){
	return &ServeMuxEx {
		http.NewServeMux(),
	}
}
