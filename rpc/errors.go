package rpc

type HandlerError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (e *HandlerError) Error() string {
	return e.Message
}
