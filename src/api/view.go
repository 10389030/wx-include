package api

import (
	"net/http"
	"sort"
	"strings"
	"crypto/sha1"
	"log"
	"fmt"
	"encoding/hex"
	"encoding/xml"
	"time"
)

func CheckServer(w http.ResponseWriter, r *http.Request) {
	r.ParseForm() // update r.Form  net/url.Values
	
	var signature = r.Form.Get("signature")
	var timestamp = r.Form.Get("timestamp")
	var nonce     = r.Form.Get("nonce")
	var echostr   = r.Form.Get("echostr")

	var fileds = []string{API_TOKEN, timestamp, nonce}
	sort.Strings(fileds)
	var txt = strings.Join(fileds, "")
	var sha1Rst = sha1.Sum([]byte(txt))
	var hexRst = hex.EncodeToString(sha1Rst[:])

	log.Printf("cal sha1: %s", hexRst)
	log.Printf("req sha1: %s", signature)

	if hexRst == signature {
		// check ok
		log.Print("check sucess!")
		fmt.Fprint(w, echostr)	
	} else {
		// check error
		log.Printf("check error: token=%s.", API_TOKEN)
		fmt.Fprint(w, "check error")
	}
}

func EventSubsribe(w http.ResponseWriter, msg *Message) {
	result := &TextMessage{
		MessageBase: MessageBase {
			ToUserName: msg.FromUserName,
			FromUserName: msg.ToUserName,
			CreateTime: time.Now().Unix(),
			MsgType: "text",
		},
		Content: "Hi! My Tel: 17097228030",
	}

	bytes, _ := xml.Marshal(result)
	log.Print(string(bytes))
	fmt.Fprint(w, bytes)
}
