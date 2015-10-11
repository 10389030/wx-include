package auth

import (
	"tool"
	"api"
)

var Handler = tool.NewServeMuxEx()
func init() {
	Handler.HandleFunc("/api", api.Handler)
}



