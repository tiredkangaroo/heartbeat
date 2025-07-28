package rpc

import (
	"fmt"
	"log/slog"
	"time"

	"github.com/valyala/fasthttp"
)

const DEFAULT_TIMEOUT = time.Minute * 1

type RequestCodec[R any, X any] interface {
	MarshalRequest(req *R, fhttp *fasthttp.Request) error
	UnmarshalRequest(fhttp *fasthttp.RequestCtx) (*R, error)

	MarshalResponse(resp *X, fhttp *fasthttp.Response) error
	UnmarshalResponse(fhttp *fasthttp.Response) (*X, error)
}

func Handler[R any, X any](codec RequestCodec[R, X], handler func(ctx *fasthttp.RequestCtx, req *R) (*X, error)) func(ctx *fasthttp.RequestCtx) {
	return func(ctx *fasthttp.RequestCtx) {
		req, err := codec.UnmarshalRequest(ctx)
		if err != nil {
			ctx.SetStatusCode(fasthttp.StatusBadRequest)
			ctx.WriteString(err.Error())
			return
		}

		resp, err := handler(ctx, req)
		if err != nil {
			e := err.(*HandlerError)
			ctx.SetStatusCode(e.Code)
			ctx.WriteString(e.Message)
			return
		}

		err = codec.MarshalResponse(resp, &ctx.Response)
		if err != nil {
			ctx.SetStatusCode(fasthttp.StatusInternalServerError)
			slog.Error("handler: marshal response", "error", err)
			return
		}
	}
}

func Perform[R any, X any](codec RequestCodec[R, X], req *R) (*X, error) {
	fhttpRequest := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(fhttpRequest)
	fhttpResponse := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(fhttpResponse)

	err := codec.MarshalRequest(req, fhttpRequest)
	if err != nil {
		return nil, fmt.Errorf("marshal request: %w", err)
	}

	err = fasthttp.DoTimeout(fhttpRequest, fhttpResponse, DEFAULT_TIMEOUT)
	if err != nil {
		return nil, fmt.Errorf("do request: %w", err)
	}

	resp, err := codec.UnmarshalResponse(fhttpResponse)
	if err != nil {
		return nil, fmt.Errorf("unmarshal response: %w", err)
	}

	return resp, nil
}
