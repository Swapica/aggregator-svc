/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

import "github.com/Swapica/aggregator-svc/internal/amount"

type CreateOrderRequest struct {
	AmountToBuy  amount.Amount `json:"amount_to_buy"`
	AmountToSell amount.Amount `json:"amount_to_sell"`
	DestChain    string        `json:"dest_chain"`
	Sender       string        `json:"sender"`
	SrcChain     string        `json:"src_chain"`
	// address of token contract
	TokenToBuy string `json:"token_to_buy"`
	// address of token contract
	TokenToSell string `json:"token_to_sell"`
}
