package rpc_plugin

type PluginOptions struct {
	SvrAddr         string   // server address
	Services        []string // service arrays
	SelectorSvrAddr string   // server discovery address ，e.g. consul server address
	TracingSvrAddr  string   // tracing server address，e.g. jaeger server address
}

// Option provides operations on Options
type PluginOption func(*PluginOptions)

// WithSvrAddr allows you to set SvrAddr of Options
func WithSvrAddr(addr string) PluginOption {
	return func(o *PluginOptions) {
		o.SvrAddr = addr
	}
}

// WithSvrAddr allows you to set Services of Options
func WithServices(services []string) PluginOption {
	return func(o *PluginOptions) {
		o.Services = services
	}
}

// WithSvrAddr allows you to set SelectorSvrAddr of Options
func WithSelectorSvrAddr(addr string) PluginOption {
	return func(o *PluginOptions) {
		o.SelectorSvrAddr = addr
	}
}

// WithSvrAddr allows you to set TracingSvrAddr of Options
func WithTracingSvrAddr(addr string) PluginOption {
	return func(o *PluginOptions) {
		o.TracingSvrAddr = addr
	}
}
