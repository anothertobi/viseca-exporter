package app

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/anothertobi/viseca-exporter/internal/csv"
	"github.com/anothertobi/viseca-exporter/pkg/viseca"
	"github.com/urfave/cli/v2"
)

const flagDateFrom = "date-from"
const flagDateTo = "date-to"
const flagForeignCurrency = "foreign-currency"

// NewTransactionsCommand creates a new transactions CLI command
func NewTransactionsCommand() *cli.Command {
	return &cli.Command{
		Name:      "transactions",
		Usage:     "list all transactions for given card id",
		ArgsUsage: "cardID",
		Flags: []cli.Flag{
			&cli.TimestampFlag{
				Name:     flagDateFrom,
				Usage:    "from which date on transactions should be fetched (format: 2006-01-02)",
				Layout:   "2006-01-02",
				Timezone: time.Local,
			},
			&cli.TimestampFlag{
				Name:     flagDateTo,
				Usage:    "to which date transactions should be fetched (format: 2006-01-02)",
				Layout:   "2006-01-02",
				Timezone: time.Local,
			},
			&cli.BoolFlag{
				Name:  flagForeignCurrency,
				Usage: "if foreign currencies and the original amount should be included in the output",
			},
		},
		Action: func(cCtx *cli.Context) error {
			return transactions(cCtx)
		},
	}
}

func transactions(cCtx *cli.Context) error {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	if !cCtx.Args().Present() {
		return errors.New("requires card id arg")
	}
	cardID := cCtx.Args().First()

	visecaClient, err := loginCLI(ctx, cCtx)
	if err != nil {
		return err
	}

	listOptions := viseca.NewDefaultListOptions()
	dateFrom := cCtx.Timestamp(flagDateFrom)
	if dateFrom != nil {
		listOptions.DateFrom = *dateFrom
	}
	dateTo := cCtx.Timestamp(flagDateTo)
	if dateTo != nil {
		listOptions.DateFrom = *dateTo
	}

	// returns all transactions if date from
	transactions, err := visecaClient.ListAllTransactionsOpts(ctx, cardID, listOptions)
	if err != nil {
		return err
	}

	includeForeignCurrency := cCtx.Bool(flagForeignCurrency)

	options := csv.TransactionOptions{
		IncludeForeignCurrency: includeForeignCurrency,
	}

	out := csv.TransactionsStringWithOptions(transactions, options)
	fmt.Print(out)

	return nil
}
