package main

import "github.com/nt-h4rd/ext-kit/transport/ws"

type Message struct {
	Type    ws.MessageType
	Payload []byte
}
