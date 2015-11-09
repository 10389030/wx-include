package auth

import (
	"api"
	"tool"
)

var Handler = tool.NewServeMuxEx()

func init() {
	Handler.HandleFunc("/api", api.Handler)
}
