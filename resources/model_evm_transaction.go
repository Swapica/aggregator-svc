/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

type EvmTransaction struct {
	Key
	Attributes    EvmTransactionAttributes    `json:"attributes"`
	Relationships EvmTransactionRelationships `json:"relationships"`
}
type EvmTransactionResponse struct {
	Data     EvmTransaction `json:"data"`
	Included Included       `json:"included"`
}

type EvmTransactionListResponse struct {
	Data     []EvmTransaction `json:"data"`
	Included Included         `json:"included"`
	Links    *Links           `json:"links"`
}

// MustEvmTransaction - returns EvmTransaction from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustEvmTransaction(key Key) *EvmTransaction {
	var evmTransaction EvmTransaction
	if c.tryFindEntry(key, &evmTransaction) {
		return &evmTransaction
	}
	return nil
}
