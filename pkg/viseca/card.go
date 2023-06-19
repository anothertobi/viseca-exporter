package viseca

import (
	"context"
	"fmt"
)

// ListTransactions returns the transactions for the given card and listOptions.
func (client *Client) ListTransactions(ctx context.Context, card string, listOptions ListOptions) (*Transactions, error) {
	path := fmt.Sprintf("card/%s/transactions", card)

	request, err := client.NewRequest(path, "GET", nil)
	if err != nil {
		return nil, err
	}
	addListOptions(request.URL, listOptions)

	transactions := &Transactions{}

	_, err = client.Do(ctx, request, transactions)
	if err != nil {
		return nil, err
	}

	return transactions, err
}

// ListAllTransactions lists all transactions for the given card.
func (client *Client) ListAllTransactions(ctx context.Context, card string) ([]Transaction, error) {
	listOptions := NewDefaultListOptions()
	return client.ListAllTransactionsOpts(ctx, card, listOptions)
}

// ListAllTransactions lists all transactions for the given card and options.
func (client *Client) ListAllTransactionsOpts(ctx context.Context, card string, listOptions ListOptions) ([]Transaction, error) {
	var transactions []Transaction

	for {
		transactionsPage, err := client.ListTransactions(ctx, card, listOptions)
		if err != nil {
			return nil, err
		}

		transactions = append(transactions, transactionsPage.Transactions...)

		listOptions.Offset += listOptions.PageSize
		if listOptions.Offset > transactionsPage.TotalCount {
			break
		}
	}

	return transactions, nil
}
