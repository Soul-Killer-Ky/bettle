package server

import (
	"flag"
	"fmt"
	"net/http"

	"github.com/pretty66/websocketproxy"
)

var (
	// command-line options:
	// websocket server endpoint
	imServerEndpoint = flag.String("im-server-endpoint", "ws://im-svc:10000/ws", "websocket server endpoint")
)

func NewWebsocketServer() http.HandlerFunc {
	wp, err := websocketproxy.NewProxy(*imServerEndpoint, func(r *http.Request) error {
		token := r.FormValue("tt")
		if token != "" {
			r.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))
		}
		// Permission to verify
		//r.Header.Set("Cookie", "----")
		// Source of disguise
		//r.Header.Set("Origin", "http://82.157.123.54:9010")
		return nil
	})
	if err != nil {
		panic(err)
	}

	return wp.Proxy
}
