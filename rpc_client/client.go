package rpc_client

import "context"

//client interface
type Client interface {
	Invoke(ctx context.Context)
}

// 全局单例client
type defaultClient struct {
	opts *ClientOptions
}

var DefaultClient = NewDefaultClient()

var NewDefaultClient = func() *defaultClient {
	return &defaultClient{
		opts: &ClientOptions{
			protocol: "proto",
		},
	}
}

func (dc *defaultClient) Invoke(ctx context.Context, req, resp interface{}, service_path string, opts ...ClientOption) error {
	for _, o := range opts {
		o(dc.opts)
	}
	return nil
}

// call 对外接口
func (dc *defaultClient) Call(ctx context.Context, req, resp interface{}, service_path string, opts ...ClientOption) error {
	err := dc.Invoke(ctx, req, resp, service_path, opts...)
	if err != nil {
		return err
	}
	return nil
}
