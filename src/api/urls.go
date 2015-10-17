package api

var Handler = NewMessageRouter()

func init() {
	Handler.HandleMsg(&MessageRoute{MsgType: "event", Event: "subscribe"}, EventSubsribe)
	Handler.HandleMsg(&MessageRoute{MsgType: "text", Event: ""}, AutoReplyText)
}
