package rpc

import (
	"encoding/hex"
	"fmt"
	"log/slog"

	"github.com/valyala/fasthttp"
)

type Server struct {
	Address  string
	handlers []fasthttp.RequestHandler
}

func (s *Server) Register(id int, handler fasthttp.RequestHandler) {
	slog.Info("registering handler", "id", idString(id))
	s.handlers[id] = handler
}

func (s *Server) ListenAndServe() error {
	return fasthttp.ListenAndServe(s.Address, func(ctx *fasthttp.RequestCtx) {
		p := ctx.Path()
		if len(p) != 3 || p[0] != '/' {
			ctx.SetStatusCode(fasthttp.StatusBadRequest)
			return
		}
		var id [1]byte
		_, err := hex.Decode(id[:], p[1:])
		if err != nil {
			ctx.SetStatusCode(fasthttp.StatusBadRequest)
			return
		}
		fmt.Println("Handling request for ID:", id[0])
		handler := s.handlers[int(id[0])]
		if handler == nil {
			ctx.SetStatusCode(fasthttp.StatusNotFound)
			return
		}
		handler(ctx)
	})
}

func NewServer(address string) *Server {
	return &Server{
		Address:  address,
		handlers: make([]fasthttp.RequestHandler, 256), // len for 256 handlers
	}
}
