package api

import (
	"tool"
)

var Handler = tool.NewServeMuxEx()
func init() {
	Handler.HandleFunc("check_svr", CheckServer)
}
