package httpBase

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "You are at %s", req.URL.Path)
	})
	http.HandleFunc("/header", func(w http.ResponseWriter, req *http.Request) {
		for k, v := range req.Header {
			fmt.Fprintf(w, "%v\t%v\n", k, v)
		}
	})
	log.Fatal(http.ListenAndServe(":8000", nil))
}
