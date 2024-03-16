package rpc_server

import (
	"context"
	"reflect"
)

type Server struct {
	opts    *ServerOptions
	service Service

	closing bool // whether the server is closing
}

func NewServer(opt ...ServerOption) *Server {
	s := &Server{
		opts: &ServerOptions{},
	}

	for _, o := range opt {
		o(s.opts)
	}
	s.service = NewService(s.opts)
	return s
}

func GetServiceMethods(p reflect.Type, v reflect.Value) ([]*MethodDesc, error) {
	var methods []*MethodDesc

	return methods, nil
}

func (s *Server) Register(service_description *ServiceDesc, svr interface{}) {
	if service_description == nil || svr == nil {
		return
	}
	ht := reflect.TypeOf(service_description.HandlerType).Elem()
	st := reflect.TypeOf(svr)

	if !st.Implements(ht) {
		// print log
	}
	ser := &service{
		svr:          svr,
		service_name: service_description.ServiceName,
		handlers:     make(map[string]Handler),
	}

	for _, method := range service_description.Methods {
		ser.handlers[method.MethodName] = method.Handler
	}

	s.service = ser
}

func (s *Server) RegisterService(service_name string, service interface{}) error {
	svr_type := reflect.TypeOf(service)
	svr_val := reflect.ValueOf(service)

	sd := &ServiceDesc{
		ServiceName: service_name,
		Svr:         service,
		HandlerType: nil,
	}

	methods, err := GetServiceMethods(svr_type, svr_val)
	if err != nil {
		return err
	}
	sd.Methods = methods
	s.Register(sd, service)
	return nil
}

func (s *Server) Handle(ctx context.Context) {
	// 超时机制
	if s.opts.time_out != 0 {
		var cancel context.CancelFunc
		ctx, cancel = context.WithTimeout(ctx, s.opts.time_out)
		defer cancel()
	}
}
