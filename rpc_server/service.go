package rpc_server

import (
	"context"
	"fmt"
)

type Handler func(ctx context.Context)

type Service interface {
	Register(string, Handler)
	Serve(options *ServerOptions)
	Close()
}

type service struct {
	ctx          context.Context
	service_name string
	handlers     map[string]Handler
	opts         *ServerOptions
	svr          interface{} //sever
	cancel       context.CancelFunc
	closing      bool
}

func (s *service) Register(handler_name string, handler Handler) {
	if s.handlers == nil {
		s.handlers = make(map[string]Handler)
	}
	s.handlers[handler_name] = handler
}

func (s *service) Serve(options *ServerOptions) {
	fmt.Println("service is serving......")
	s.opts = options
}

func NewService(opts *ServerOptions) Service {
	return &service{
		opts: opts,
	}
}

func (s *service) Close() {
	s.closing = true
	if s.cancel != nil {
		s.cancel()
	}
	fmt.Println("service is closing ........")
}

func (s *service) Name() string {
	return s.service_name
}

func (s *service) Handle(ctx context.Context) {
	fmt.Println("service handling ..........")
}
