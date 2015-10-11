package auth

import (
    "net/http"
	"fmt"

	"tool"
)

func auth_svr(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello go server!")
}


var Handler = tool.NewServeMuxEx()
func init() {
	Handler.HandleFunc("/abc", auth_svr)
}



