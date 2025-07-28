package main

import (
	"fmt"
	"log/slog"

	"github.com/tiredkangaroo/heartbeat/rpc"
	"github.com/valyala/fasthttp"
)

func main() {
	srv := rpc.NewServer(":8000")
	srv.Register(rpc.HandlerHeartbeat, rpc.Handler(rpc.DefaultHeartbeatCodec, func(ctx *fasthttp.RequestCtx, req *rpc.RequestHeartbeat) (*rpc.ResponseHeartbeat, error) {
		return &rpc.ResponseHeartbeat{OK: true}, nil
	}))

	fmt.Println("server listening on", srv.Address)
	if err := srv.ListenAndServe(); err != nil {
		slog.Error("listen and serve (fatal)", "error", err)
		return
	}
}
