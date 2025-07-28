package rpc

import (
	"encoding/hex"
)

const (
	HandlerHeartbeat = iota
)

func idString(id int) string {
	switch id {
	case HandlerHeartbeat:
		return "heartbeat"
	}
	return "unknown"
}

type RequestHeartbeat struct{}
type ResponseHeartbeat struct {
	OK bool `json:"ok"`
}

func pathFor(id int) string {
	var pathBytes = [3]byte{
		'/',
	}
	hex.Encode(pathBytes[1:], []byte{byte(id)})
	return string(pathBytes[:])
}
