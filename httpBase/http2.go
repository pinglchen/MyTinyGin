package httpBase

import (
	"fmt"
	"net/http"
)

type Engine struct{}

func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	fmt.Println("request...")
	switch req.URL.Path {
	case "/":
		fmt.Fprintf(w, "You are at %s", req.URL.Path)
	case "/header":
		for k, v := range req.Header {
			fmt.Fprintf(w, "%v\t%v\n", k, v)
		}
	default:
		fmt.Fprint(w, "404 NOT FOUND")
	}
}
