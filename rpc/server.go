package rpc

import "github.com/valyala/fasthttp"

type Server struct {
	Address  string
	handlers []fasthttp.RequestHandler
}

func (s *Server) Register(id int, handler fasthttp.RequestHandler) {
	s.handlers[id] = handler
}

func (s *Server) ListenAndServe() error {
	return fasthttp.ListenAndServe(s.Address, func(ctx *fasthttp.RequestCtx) {
		p := ctx.Path()
		if len(p) < 1 {
			ctx.SetStatusCode(fasthttp.StatusBadRequest)
			ctx.WriteString("invalid path")
			return
		}
		handler := s.handlers[int(p[0])]
		handler(ctx)
	})
}

func NewServer(address string) *Server {
	return &Server{
		Address:  address,
		handlers: make([]fasthttp.RequestHandler, 256), // len for 256 handlers
	}
}
