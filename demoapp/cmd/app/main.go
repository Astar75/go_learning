package main

import (
	"demoapp/demoapp/pkg/stringutils"
	"net/http"
)

func main() {
	http.ListenAndServe(
		":8080",
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			s := r.URL.Query()["s"]
			rev := stringutils.Rev(s[0])
			w.Write([]byte(rev))
		},
		),
	)
}
