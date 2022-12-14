/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

type CancelOrderRequest struct {
	// The identifier of the order that was selected by the user to cancel.
	OrderId int `json:"order_id"`
	// The address of the sender
	Sender string `json:"sender"`
	// The identifier of the chain in which the order is placed with which the match occurred.
	SrcChain string `json:"src_chain"`
}
