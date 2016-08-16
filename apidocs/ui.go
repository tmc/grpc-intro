package apidocs

import (
	"net/http"

	"github.com/GeertJohan/go.rice"
	"github.com/tmc/grpc-intro/apidocs/specs"
)

var (
	URLPrefixUI = "/apidocs/"
	HandlerUI   = http.StripPrefix(URLPrefixUI, http.FileServer(rice.MustFindBox("../third_party/swagger-ui/").HTTPBox()))

	URLPrefix = "/apidocs-specs/"
	Handler   = CORS(http.StripPrefix(URLPrefix, http.FileServer(specs.Specifications.HTTPBox())))
)

func init() {
	rice.Debug = true
}

func CORS(h http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", r.Header.Get("Access-Control-Request-Headers"))
		if r.Method == "OPTIONS" {
			return
		}
		h.ServeHTTP(w, r)
	}
}
