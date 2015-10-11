package tool

import (
	"net/http"
	"log"
)

// Extern net/http.ServeMux to add subtree rooting.
// according to net/http.pathMatch, only when a pattern is ending with '/' while be
// regard as subtree rooting.
// ATTENTION: r.URL.PATH may be changed. You cannot get the actually PATH anymore.

type ServeMuxEx struct {
	*http.ServeMux
}

func (mux *ServeMuxEx) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.RequestURI == "*" {
		if r.ProtoAtLeast(1, 1) {
			w.Header().Set("Connection", "close")
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	h, pattern := mux.Handler(r)
	log.Printf("pattern: %s, url=%s", pattern, r.URL.Path)
	var patternLen = len(pattern)
	if patternLen > 0 && pattern[patternLen - 1] == '/' { 
		r.URL.Path = r.URL.Path[len(pattern) - 1:]
	}

	h.ServeHTTP(w, r)
}

func NewServeMuxEx() (*ServeMuxEx){
	return &ServeMuxEx {
		http.NewServeMux(),
	}
}
