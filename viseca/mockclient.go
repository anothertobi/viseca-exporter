package viseca

import (
	"context"

	"github.com/stretchr/testify/mock"
)

type MockedVisecaAPI struct {
	mock.Mock
}

func (m *MockedVisecaAPI) ListAllTransactions(ctx context.Context, card string, listOptions ListOptions) (*Transactions, error) {
	return nil, nil
}
