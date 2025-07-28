package rpc

import "github.com/valyala/fasthttp"

type HeartbeatCodec struct{}

func (c *HeartbeatCodec) MarshalRequest(req *RequestHeartbeat, fhttp *fasthttp.Request) error {
	fhttp.URI().SetPath(pathFor(HandlerHeartbeat))
	return nil
}
func (c *HeartbeatCodec) UnmarshalRequest(fhttp *fasthttp.RequestCtx) (*RequestHeartbeat, error) {
	return &RequestHeartbeat{}, nil
}
func (c *HeartbeatCodec) MarshalResponse(resp *ResponseHeartbeat, fhttp *fasthttp.Response) error {
	if resp.OK {
		fhttp.SetStatusCode(fasthttp.StatusOK)
	} else {
		fhttp.SetStatusCode(fasthttp.StatusInternalServerError)
	}
	return nil
}
func (c *HeartbeatCodec) UnmarshalResponse(fhttp *fasthttp.Response) (*ResponseHeartbeat, error) {
	resp := &ResponseHeartbeat{}
	if fhttp.StatusCode() == fasthttp.StatusOK {
		resp.OK = true
	} else {
		resp.OK = false
	}
	return resp, nil
}

var DefaultHeartbeatCodec = &HeartbeatCodec{}
