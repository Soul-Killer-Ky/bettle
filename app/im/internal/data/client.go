package data

import "github.com/gorilla/websocket"

type imClient struct {
	ID     uint
	Socket *websocket.Conn
}
