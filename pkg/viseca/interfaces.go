package viseca

import "context"

// VisecaAPI provides an interface to a subset of the Viseca API
type VisecaAPI interface {
	// ListTransactions returns the transactions for the given card.
	ListTransactions(ctx context.Context, card string, listOptions ListOptions) (*Transactions, error)
}

var _ VisecaAPI = (*Client)(nil)
