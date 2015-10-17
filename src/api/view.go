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

func ReplyText(w http.ResponseWriter, msg *Message, content string) {
	result := &TextMessage {
		MessageBase: MessageBase {
			ToUserName: msg.FromUserName,
			FromUserName: msg.ToUserName,
			CreateTime: time.Now().Unix(),
			MsgType: "text",
		},
	}

	str_slice := strings.Split(content, "\n")
	for idx, str := range str_slice {
		str_slice[idx] = strings.TrimSpace(str)
	}
	result.Content = strings.Join(str_slice, "\n")

	data, _ := xml.Marshal(result)
	fmt.Fprint(w, string(data))

	log.Printf("return: %s", string(data))
}

// toolkit: return a handler witch return specified text
func _wrap(content string) (f func(http.ResponseWriter, *Message)) {
	return func(w http.ResponseWriter, msg *Message) {
		ReplyText(w, msg, content)
	}
}


func EventSubsribe(w http.ResponseWriter, msg *Message) {
	log.Printf("msg: %#v", msg)

	ReplyText(w, msg, "朋来,乐哉!\n即课满足你一个愿望，大胆说出来吧... (限本人力所能及) \n例如: wish 大帅哥/萌妹子一个")
}

func AutoReplyText(w http.ResponseWriter, msg *Message) {

	log.Print("msg: %#v", msg)

	type Command struct {
		Note string	
		Handler func(w http.ResponseWriter, msg *Message)
	}

	// cmd must be lower
	// : please register cmd into cmds
	cmds := map[string] *Command {
		"about": &Command{
			Note: "本订阅号基本信息",
			Handler: _wrap(`
				Author   : Junzexu
				Version  : V0.0.1_1 
				Mail     : xu_jun_wei@126.com
				Power By : WXSVR / AWS / GoLang
			`),
		},
		"wish": &Command{
			Note: "我是阿拉神灯",
			Handler: _wrap("骗你的啦！\n听老板说好好工作，想要的都会有的，哈哈哈..."),
		},
	}

	cmd := strings.ToLower(msg.Content)
	if command, ok := cmds[cmd]; ok {
		command.Handler(w, msg);
	} else {
		cmds["cmd"].Handler(w, msg)
	}


	// attention: following code should alway put at the end
	cmds["cmd"] = &Command {
		Note: "列出所有的指令",
		Handler: nil,
	}

	max_length := 0
	for k, _ := range cmds {
		chars_len := len([]rune(k))
		if max_length < chars_len {
			max_length = chars_len
		}
	}

	segs := []string{}
	for k, v := range cmds {
		segs = append(segs, k + ": \t" + v.Note)
	}

	commands_txt := strings.Join(segs, "\n")
	cmds["cmd"].Handler = _wrap(commands_txt)
}
