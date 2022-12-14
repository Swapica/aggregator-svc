package mem

import "github.com/Swapica/aggregator-svc/internal/data"

func NewNodesQ(nodes []string) data.NodesQ {
	return &nodesQ{
		nodes: nodes,
	}
}

type nodesQ struct {
	nodes []string
}

func (q *nodesQ) New() data.NodesQ {
	return NewNodesQ(q.nodes)
}

func (q *nodesQ) Select() ([]string, error) {
	return q.nodes, nil
}
