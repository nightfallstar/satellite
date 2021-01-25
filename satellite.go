package main

import (
	"fmt"
	"log"
	"github.com/valyala/fasthttp"
)

func httpsHandler(ctx *fasthttp.RequestCtx) {
	switch string(ctx.Path()) {
	case "/auth":
		// TO DO
		break

	default:
		fmt.Fprintf(ctx, "{}")
	}
}

func main() {
	log.Fatal(fasthttp.ListenAndServeTLS("0.0.0.0:8443", "certs/development/cert.pem", "certs/development/key.pem", httpsHandler))
}
