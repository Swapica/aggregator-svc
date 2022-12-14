package config

import (
	"gitlab.com/distributed_lab/figure/v3"
	"gitlab.com/distributed_lab/kit/comfig"
	"gitlab.com/distributed_lab/kit/kv"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

func NewNoder(getter kv.Getter) Noder {
	return &noder{
		getter: getter,
	}
}

type Noder interface {
	Nodes() []string
}

type noder struct {
	getter kv.Getter
	once   comfig.Once

	nodes []string
}

func (c *noder) Nodes() []string {
	c.readConfig()
	return c.nodes
}

func (c *noder) readConfig() {
	c.once.Do(func() interface{} {
		cfg := struct {
			Nodes []string `fig:"nodes,required"`
		}{}
		err := figure.
			Out(&cfg).
			From(kv.MustGetStringMap(c.getter, "data")).
			Please()
		if err != nil {
			panic(errors.Wrap(err, "failed to figure out nodes"))
		}

		for i, node := range cfg.Nodes {
			cfg.Nodes[i] = node
		}

		c.nodes = cfg.Nodes

		return nil
	})
}
