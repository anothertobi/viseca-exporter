package viseca

import (
	"context"
)

// GetUser returns the user information.
func (client *Client) GetUser(ctx context.Context) (*User, error) {
	request, err := client.NewRequest("user", "GET", nil)
	if err != nil {
		return nil, err
	}

	user := &User{}

	_, err = client.Do(ctx, request, user)
	if err != nil {
		return nil, err
	}

	return user, err
}
