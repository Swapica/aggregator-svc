package models

import (
	"github.com/Swapica/aggregator-svc/internal/data"
	"github.com/Swapica/swapica-svc/resources"
)

func newChainKey(id string) resources.Key {
	return resources.Key{
		ID:   id,
		Type: resources.CHAIN,
	}
}

func newChainModel(value data.Chain) *resources.Chain {
	return &resources.Chain{
		Key: newChainKey(value.ID),
		Attributes: resources.ChainAttributes{
			Name:        value.Name,
			Icon:        value.Icon,
			ChainType:   value.Type,
			ChainParams: value.ChainParams,
		},
	}
}

func newChainModelWithRelation(value data.Chain) *resources.Chain {
	model := newChainModel(value)

	return model
}

func NewChainListResponse(chains []data.Chain) resources.ChainListResponse {
	response := resources.ChainListResponse{
		Data:     make([]resources.Chain, len(chains)),
		Included: resources.Included{},
	}

	for i, chain := range chains {
		response.Data[i] = *newChainModelWithRelation(chain)
	}

	return response
}
