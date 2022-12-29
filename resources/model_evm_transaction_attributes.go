/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

type EvmTransactionAttributes struct {
	// Is enough signatures to send tx
	Confirmed *bool                `json:"confirmed,omitempty"`
	TxBody    EvmTransactionTxBody `json:"tx_body"`
}
