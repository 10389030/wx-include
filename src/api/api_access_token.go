package api

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

type AccessTokenRsp struct {
	AccessToken string `json:"access_token"`
	ExpireIn    int64  `json:"expires_in"`
	ErrCode     int    `json:"errcode"`
	ErrMsg      string `json:"errmsg"`
}

func GetAccessTokenReq() (*http.Request, error) {
	var vals = url.Values{}
	vals.Set("grant_type", "client_credential")
	vals.Set("appid", API_APPID)
	vals.Set("secret", API_SECRET)

	var url = API_HOST + "/cgi-bin/token?" + vals.Encode()
	return http.NewRequest("GET", url, nil)
}

func GetAccessToken() *AccessTokenRsp {
	var cli = &http.Client{}
	var req, _ = GetAccessTokenReq()

	var rsp, err = cli.Do(req)
	if rsp != nil {
		defer rsp.Body.Close()
	}

	if err != nil {
		log.Panic("GetAccessTokenReq err: %d", err)
	}

	var rsp_st = &AccessTokenRsp{ErrCode: 200}
	var data, _ = ioutil.ReadAll(rsp.Body)
	json.Unmarshal(data, rsp_st)

	if rsp_st.ErrCode != 200 {
		log.Printf("GetAccessTokenRsp: %#v", rsp_st)
	}

	return rsp_st
}
