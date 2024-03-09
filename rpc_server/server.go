package rpc_server

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
