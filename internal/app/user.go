package app

import (
	"context"
	"encoding/json"
	"os"
	"time"

	"github.com/urfave/cli/v2"
)

// NewUserCommand creates a new user CLI command
func NewUserCommand() *cli.Command {
	return &cli.Command{
		Name:  "user",
		Usage: "get user",
		Action: func(cCtx *cli.Context) error {
			return user(cCtx)
		},
	}
}

func user(cCtx *cli.Context) error {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	visecaClient, err := loginCLI(ctx, cCtx)
	if err != nil {
		return err
	}

	user, err := visecaClient.GetUser(ctx)
	if err != nil {
		return err
	}

	encoder := json.NewEncoder(os.Stdout)

	encoder.Encode(user)

	return nil
}
