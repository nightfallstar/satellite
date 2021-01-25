package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func (res http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(res, "{}")
	})

	log.Fatal(http.ListenAndServeTLS("0.0.0.0:8443", "certs/development/cert.pem", "certs/development/key.pem", nil))
}
