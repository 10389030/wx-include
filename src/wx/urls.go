package main

import (
	"api"
	"net/http"
)

var Handler = http.NewServeMux()

func init() {
	Handler.Handle("/api/", api.Handler)
}

