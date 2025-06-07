package csv

import (
	"fmt"
	"strings"

	"github.com/anothertobi/viseca-exporter/pkg/viseca"
)

type TransactionOptions struct {
	IncludeForeignCurrency bool
}

// TransactionsString returns a CSV representation of the transactions according to the provided options.
func TransactionsStringWithOptions(transactions []viseca.Transaction, options TransactionOptions) string {
	var stringBuilder strings.Builder

	stringBuilder.WriteString(`"TransactionID","Date","Merchant","Amount",`)
	if options.IncludeForeignCurrency {
		stringBuilder.WriteString(`"Currency","OriginalAmount","OriginalCurrency",`)
	}
	stringBuilder.WriteString(`"PFMCategoryID","PFMCategoryName"`)
	stringBuilder.WriteString("\n")

	for _, transaction := range transactions {
		stringBuilder.WriteString(TransactionStringWithOptions(transaction, options))
		stringBuilder.WriteString("\n")
	}

	return stringBuilder.String()
}

// TransactionsString returns a CSV representation of the transactions.
func TransactionsString(transactions []viseca.Transaction) string {
	options := TransactionOptions{
		IncludeForeignCurrency: false,
	}
	return TransactionsStringWithOptions(transactions, options)
}

// TransactionStringWithOptions returns a CSV record according to the provided options.
func TransactionStringWithOptions(transaction viseca.Transaction, options TransactionOptions) string {
	innerRecord := strings.Join([]string{
		transaction.TransactionID,
		transaction.Date,
		prettiestMerchantName(transaction),
		fmt.Sprintf("%.2f", transaction.Amount),
	}, `","`)

	if options.IncludeForeignCurrency {
		innerRecord = strings.Join([]string{
			innerRecord,
			transaction.Currency,
			fmt.Sprintf("%.2f", transaction.OriginalAmount),
			transaction.OriginalCurrency,
		}, `","`)
	}

	innerRecord = strings.Join([]string{
		innerRecord,
		transaction.PFMCategory.ID,
		transaction.PFMCategory.Name,
	}, `","`)

	return fmt.Sprintf(`"%s"`, innerRecord)
}

// TransactionString returns a CSV record.
func TransactionString(transaction viseca.Transaction) string {
	options := TransactionOptions{
		IncludeForeignCurrency: false,
	}
	return TransactionStringWithOptions(transaction, options)
}

// prettiestMerchantName extracts the prettiest merchant name from a transaction.
func prettiestMerchantName(transaction viseca.Transaction) string {
	if transaction.PrettyName != "" {
		return transaction.PrettyName
	} else {
		return transaction.MerchantName
	}
}
