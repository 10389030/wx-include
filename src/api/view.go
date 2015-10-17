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
	log.Printf("msg: %#v", msg)

	result := &TextMessage{
		MessageBase: MessageBase {
			ToUserName: msg.FromUserName,
			FromUserName: msg.ToUserName,
			CreateTime: time.Now().Unix(),
			MsgType: "text",
		},
		Content: `<!CDATA[
			Hi, 欢迎关注我的订阅号！来，机会只有一次，大声说出你的愿望吧...
		]]`,
	}

	bytes, _ := xml.Marshal(result)
	log.Print(string(bytes))
	fmt.Fprint(w, string(bytes))
}


func AutoReplyText(w http.ResponseWriter, msg *Message) {
	log.Print("msg: %#v", msg)

	result := &TextMessage{
		MessageBase: MessageBase {
			ToUserName: msg.FromUserName,
			FromUserName: msg.ToUserName,
			CreateTime: time.Now().Unix(),
			MsgType: "text",
		},
		Content: "",
	}

	var content = ""
	switch msg.Content {
		case "cmd":
			content = `
				cmd   : 获得所有指令 
				About : 订阅号信息
			`
		case "About":
			content = `
				ID       : gh_95b6702312f1
			    Author   : 许俊伟(Junzexu);
				Version  : v0.0.1_1;
				CreateAt : 2015-10-17;
				Mail     : xu_jun_wei@126.com;
			`
		default:
			content = `
				Sorry! unrecognize command.
				Try 'cmd' for all command.
			`
	}

	result.Content = "<!CDATA[" + content + "]]"

	log.Print("replay: %#v", result)
	bytes, _ := xml.Marshal(result)
	fmt.Fprint(w, string(bytes))
}
