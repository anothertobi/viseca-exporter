package csv

import (
	"fmt"
	"strings"

	"github.com/anothertobi/viseca-exporter/pkg/viseca"
)

// TransactionsString() returns a CSV representation of the transactions.
func TransactionsString(transactions []viseca.Transaction) string {
	var stringBuilder strings.Builder

	stringBuilder.WriteString(`"TransactionID","Date","Merchant","Amount","PFMCategoryID","PFMCategoryName"`)
	stringBuilder.WriteString("\n")

	for _, transaction := range transactions {
		stringBuilder.WriteString(TransactionString(transaction))
		stringBuilder.WriteString("\n")
	}

	return stringBuilder.String()
}

// TransactionString returns a CSV record.
func TransactionString(transaction viseca.Transaction) string {
	innerRecord := strings.Join([]string{
		transaction.TransactionID,
		transaction.Date,
		prettiestMerchantName(transaction),
		fmt.Sprintf("%.2f", transaction.Amount),
		transaction.PFMCategory.ID,
		transaction.PFMCategory.Name,
	}, `","`)

	return fmt.Sprintf(`"%s"`, innerRecord)
}

// prettiestMerchantName extracts the prettiest merchant name from a transaction.
func prettiestMerchantName(transaction viseca.Transaction) string {
	if transaction.PrettyName != "" {
		return transaction.PrettyName
	} else {
		return transaction.MerchantName
	}
}
