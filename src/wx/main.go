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
	if rsp.Errcode == 200 {
		log.Printf("%#v", rsp.Data)
		log.Print("access_token: %s", (*rsp.Data)["access_token"].(string))
		log.Print("expires_in: %v", (*rsp.Data)["expires_in"].(float64))
	} else {
		log.Print("GetAccessToken error")
	}

	log.Fatal(svr.ListenAndServe())
}
