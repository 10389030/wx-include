package api

import (
	"net/http"
	"net/url"
	"encoding/json"
	"log"
	"io/ioutil"
)

type ApiData struct {
	Errcode int
	Data    *map[string]interface {}
}

func GetAccessTokenReq() (*http.Request, error) {
	var vals = url.Values{}
	vals.Set("grant_type", "client_credential")
	vals.Set("appid", API_APPID)
	vals.Set("secret", API_SECRET)

	var url = API_HOST + "/cgi-bin/token?" + vals.Encode() 
	return http.NewRequest("GET", url, nil)
}

func GetAccessToken() (rst *ApiData){
	var cli = &http.Client{}
	var req, ok = GetAccessTokenReq()
	if ok != nil {
		log.Panic("GetAccessTokenReq: %#v", ok)
	}

	var rsp, _ = cli.Do(req)
	defer rsp.Body.Close()

	rst = &ApiData{Errcode: 200, Data: &map[string]interface{}{}}
	var data, _ = ioutil.ReadAll(rsp.Body)
	json.Unmarshal(data, rst.Data)
	
	if rsp.StatusCode != 200 {
		rst.Errcode = (*rst.Data)["errcode"].(int)
	}

	return rst
}
