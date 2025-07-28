package main

import (
	"log/slog"

	"github.com/tiredkangaroo/heartbeat/rpc"
	"github.com/valyala/fasthttp"
)

func main() {
	srv := rpc.NewServer(":8080")
	srv.Register(rpc.HandlerHeartbeat, rpc.Handler(rpc.DefaultHeartbeatCodec, func(ctx *fasthttp.RequestCtx, req *rpc.RequestHeartbeat) (*rpc.ResponseHeartbeat, error) {
		return &rpc.ResponseHeartbeat{OK: true}, nil
	}))

	if err := srv.ListenAndServe(); err != nil {
		slog.Error("listen and serve (fatal)", "error", err)
		return
	}
}
