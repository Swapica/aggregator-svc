/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

type CreateMatchRequest struct {
	// The number of tokens that the user provided to purchase  tokens that were placed in the order
	AmountToSell int `json:"amount_to_sell"`
	// The identifier of the chain in which the match occurred to a particular order.
	DestChain string `json:"dest_chain"`
	// The identifier of the order that was selected by the user to create the match.
	OrderId int `json:"order_id"`
	// redeem tx data with collected signatures
	RawTxData *string `json:"raw_tx_data,omitempty"`
	// The address of the sender
	Sender string `json:"sender"`
	// The identifier of the chain in which the order was placed.
	SrcChain string `json:"src_chain"`
	// The address of the token that the matcher wanted to sell in order  to buy the token that was offered for sale.
	TokenToSell string `json:"token_to_sell"`
}
