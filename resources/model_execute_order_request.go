/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

type ExecuteOrderRequest struct {
	// The identifier of the chain in which the order was placed.
	DestChain string `json:"dest_chain"`
	// The identifier of the match that was selected by the user.
	MatchId int `json:"match_id"`
	// The identifier of the order that was selected by the user to create the match.
	OrderId int `json:"order_id"`
	// redeem tx data with collected signatures
	RawTxData *string `json:"raw_tx_data,omitempty"`
	// The address of the receiver who will get tokens placed in order.
	Receiver string `json:"receiver"`
	// The address of the sender
	Sender string `json:"sender"`
	// The identifier of the chain in which the order is placed with which the match occurred.
	SrcChain string `json:"src_chain"`
}
