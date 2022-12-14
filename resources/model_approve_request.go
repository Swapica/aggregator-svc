/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

type ApproveRequest struct {
	// The identifier of the chain in which the token exists.
	ChainId string `json:"chain_id"`
	// The address of the sender
	Sender string `json:"sender"`
	// The token address on specified chain id.
	TokenAddress string `json:"token_address"`
	// The token address on specified chain id.
	TokenType string `json:"token_type"`
}
