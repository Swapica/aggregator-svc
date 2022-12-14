/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

type CancelMatchRequest struct {
	// The identifier of the chain in which the match was placed.
	DestChain string `json:"dest_chain"`
	// The identifier of the match that was selected by the user to cancel it.
	MatchId int `json:"match_id"`
	// redeem tx data with collected signatures
	RawTxData *string `json:"raw_tx_data,omitempty"`
	// The address of the sender
	Sender string `json:"sender"`
	// The identifier of the chain in which the order is placed with which the match occurred.
	SrcChain string `json:"src_chain"`
}
