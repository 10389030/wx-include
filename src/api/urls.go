package api

var Handler = NewMessageRouter()

func init() {
	Handler.HandleMsg(&MessageRoute{MsgType: "event", Event: "subscribe"}, EventSubsribe)
}
