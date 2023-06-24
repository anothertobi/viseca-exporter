package app

import (
	"context"
	"encoding/json"
	"os"
	"time"

	"github.com/anothertobi/viseca-exporter/pkg/viseca"
	"github.com/urfave/cli/v2"
)

// NewCardsCommand creates a new cards CLI command
func NewCardsCommand() *cli.Command {
	return &cli.Command{
		Name:  "cards",
		Usage: "list all cards",
		Action: func(cCtx *cli.Context) error {
			return cards(cCtx)
		},
	}
}

func cards(cCtx *cli.Context) error {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	visecaClient, err := loginCLI(ctx, cCtx)
	if err != nil {
		return err
	}

	cardListOptions := viseca.NewDefaultCardListOptions()
	cards, err := visecaClient.ListCards(ctx, cardListOptions)
	if err != nil {
		return err
	}

	encoder := json.NewEncoder(os.Stdout)

	encoder.Encode(cards)

	return nil
}
