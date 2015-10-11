package api

import (
	"net/http"
	"sort"
	"strings"
	"crypto/sha1"
	"log"
	"fmt"
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

	log.Print("cal sha1: % x", sha1Rst)
	log.Print("req sha1: % x", signature)

	if string(sha1Rst[:]) == signature {
		// check ok
		log.Print("check sucess!")
		fmt.Fprint(w, echostr)	
	} else {
		// check error
		log.Print("check error: token=%s.", API_TOKEN)
		fmt.Fprint(w, "check error")
	}
}
