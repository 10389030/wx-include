package api

var Handler = NewServeMuxEx()
func init() {
	Handler.HandleMsg(&MsgRouteInfo{MsgType: "event", Event: "subscribe"}, EventSubsribe)
}
