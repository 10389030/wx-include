package main

import (
	"api"
)

var Handler = api.NewServeMuxEx()

func init() {
	Handler.Handle("/api/", api.Handler)
}

