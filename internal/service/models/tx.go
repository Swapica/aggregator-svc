package models

import (
	"github.com/Swapica/aggregator-svc/internal/data"
	"github.com/Swapica/swapica-svc/resources"
)

type TxResponse struct {
	Data     interface{}        `json:"data"`
	Included resources.Included `json:"included"`
}

func NewTxResponse(tx interface{}, chain data.Chain) TxResponse {
	response := TxResponse{
		Data:     tx,
		Included: resources.Included{},
	}

	response.Included.Add(newChainModel(chain))

	return response
}
