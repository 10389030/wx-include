package main

import (
	"net/http"
	"log"
	"api"
)

func main() {
	var svr = &http.Server{
		Addr: ":http",
		Handler: Handler,
	}

	var rsp = api.GetAccessToken()
	log.Printf("%#v", rsp)

	log.Fatal(svr.ListenAndServe())
}
