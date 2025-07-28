package rpc

const (
	HandlerHeartbeat = iota
)

type RequestHeartbeat struct{}
type ResponseHeartbeat struct {
	OK bool `json:"ok"`
}
