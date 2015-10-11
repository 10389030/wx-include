package main

import (
	"tool"
	"api"
)

var Handler = tool.NewServeMuxEx()

func init() {
	Handler.Handle("/api/", api.Handler)
}

