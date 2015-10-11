package main

import (
	"tool"
	"auth"
)

var Handler = tool.NewServeMuxEx()

func init() {
	Handler.Handle("/auth/", auth.Handler)
}

