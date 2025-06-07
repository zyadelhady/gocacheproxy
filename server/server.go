package server

import "gocacheproxy/ctx"

type Server struct {
	url string
	ctx *ctx.Context
}

func New(url string,ctx *ctx.Context) *Server {
	return &Server{
		url: url,
		ctx: ctx,

	}
}
