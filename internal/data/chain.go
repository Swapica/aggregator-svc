package data

import (
	"encoding/json"
	"github.com/Swapica/swapica-svc/resources"
)

type Chain struct {
	ID          string              `fig:"id,required"`
	Name        string              `fig:"name,required"`
	Icon        *string             `fig:"icon"`
	Type        resources.ChainType `fig:"type,required"`
	ChainParams json.RawMessage     `fig:"chain_params"`

	Confirmations int    `fig:"confirmations,required"`
	SwapContract  string `fig:"swap_contract,required"`
	RpcEndpoint   string `fig:"rpc_endpoint,required"`
}
