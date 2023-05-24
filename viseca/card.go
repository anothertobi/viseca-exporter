package viseca

import (
	"context"
	"fmt"
)

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
