package rpc_plugin

type Plugin interface {
	Init(...PluginOption) error
}

// ResolverPlugin defines the standard for all server discovery plug-ins
type ResolverPlugin interface {
	Init(...PluginOption) error
}

// TracingPlugin defines the standard for all tracing plug-ins
type TracingPlugin interface {
	// todo
	//Init(...Option) (opentracing.Tracer, error)
}

// PluginMap defines a global plug-in map
var PluginMap = make(map[string]Plugin)

// Register opens an entry point for all plug-ins to register
func Register(name string, plugin Plugin) {
	if PluginMap == nil {
		PluginMap = make(map[string]Plugin)
	}
	PluginMap[name] = plugin
}
