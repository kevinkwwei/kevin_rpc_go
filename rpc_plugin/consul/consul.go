package consul

import (
	"fmt"
	"github.com/hashicorp/consul/api"
	"kevin_rpc_go/rpc_load_balancer"
	"kevin_rpc_go/rpc_plugin"
	"net/http"
	"strings"
)

const Name = "consul"

type Consul struct {
	opts         *rpc_plugin.PluginOptions
	client       *api.Client
	config       *api.Config
	balancerName string // load balancing mode, including random, polling, weighted polling, consistent hash, etc
	writeOptions *api.WriteOptions
	queryOptions *api.QueryOptions
}

func init() {
	rpc_plugin.Register(Name, ConsulSvr)
}

var ConsulSvr = &Consul{
	opts: &rpc_plugin.PluginOptions{},
}

func Init(consulSvrAddr string, opts ...rpc_plugin.PluginOption) error {
	for _, o := range opts {
		o(ConsulSvr.opts)
	}

	ConsulSvr.opts.SelectorSvrAddr = consulSvrAddr
	err := ConsulSvr.InitConfig()
	return err
}

func (c *Consul) InitConfig() error {

	config := api.DefaultConfig()
	c.config = config

	config.HttpClient = http.DefaultClient
	config.Address = c.opts.SelectorSvrAddr
	config.Scheme = "http"

	client, err := api.NewClient(config)
	if err != nil {
		return err
	}

	c.client = client

	return nil
}

func (c *Consul) Resolve(serviceName string) ([]*rpc_load_balancer.Node, error) {

	pairs, _, err := c.client.KV().List(serviceName, nil)
	if err != nil {
		return nil, err
	}

	if len(pairs) == 0 {
		return nil, fmt.Errorf("no services find in path : %s", serviceName)
	}
	var nodes []*rpc_load_balancer.Node
	for _, pair := range pairs {
		nodes = append(nodes, &rpc_load_balancer.Node{
			Key:   pair.Key,
			Value: pair.Value,
		})
	}
	return nodes, nil
}

func (c *Consul) Select(serviceName string) (string, error) {

	nodes, err := c.Resolve(serviceName)

	if nodes == nil || len(nodes) == 0 || err != nil {
		return "", err
	}

	balancer := rpc_load_balancer.GetBalancer(c.balancerName)
	node := balancer.Balance(serviceName, nodes)

	if node == nil {
		return "", fmt.Errorf("no services find in %s", serviceName)
	}

	return parseAddrFromNode(node)
}

func parseAddrFromNode(node *rpc_load_balancer.Node) (string, error) {
	if node.Key == "" {
		return "", errors.New("addr is empty")
	}

	strs := strings.Split(node.Key, "/")

	return strs[len(strs)-1], nil
}

func (c *Consul) Init(opts ...rpc_plugin.PluginOption) error {

	for _, o := range opts {
		o(c.opts)
	}

	if len(c.opts.Services) == 0 || c.opts.SvrAddr == "" || c.opts.SelectorSvrAddr == "" {
		return fmt.Errorf("consul init error, len(services) : %d, svrAddr : %s, selectorSvrAddr : %s",
			len(c.opts.Services), c.opts.SvrAddr, c.opts.SelectorSvrAddr)
	}

	if err := c.InitConfig(); err != nil {
		return err
	}

	for _, serviceName := range c.opts.Services {
		nodeName := fmt.Sprintf("%s/%s", serviceName, c.opts.SvrAddr)

		kvPair := &api.KVPair{
			Key:   nodeName,
			Value: []byte(c.opts.SvrAddr),
			Flags: api.LockFlagValue,
		}

		if _, err := c.client.KV().Put(kvPair, c.writeOptions); err != nil {
			return err
		}
	}

	return nil
}
