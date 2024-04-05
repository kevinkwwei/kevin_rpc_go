package rpc_load_balancer

type Node struct {
	Key    string
	Value  []byte
	weight int
}

type Balancer interface {
	Balance(string, []*Node) *Node
}

var balancerMap = make(map[string]Balancer, 0)

const (
	Random             = "random"
	RoundRobin         = "roundRobin"
	WeightedRoundRobin = "weightedRoundRobin"
	ConsistentHash     = "consistentHash"

	Custom = "custom"
)

func init() {
	RegisterBalancer(Random, DefaultBalancer)
	RegisterBalancer(RoundRobin, RRBalancer)
	RegisterBalancer(WeightedRoundRobin, WRRBalancer)
}

func RegisterBalancer(name string, balancer Balancer) {
	if balancerMap == nil {
		balancerMap = make(map[string]Balancer)
	}
	balancerMap[name] = balancer
}

func GetBalancer(name string) Balancer {
	if balancer, ok := balancerMap[name]; ok {
		return balancer
	}
	return DefaultBalancer
}

var DefaultBalancer = newRandomBalancer()

var RRBalancer = newRoundRobinBalancer()

var WRRBalancer = newWeightedRoundRobinBalancer()
