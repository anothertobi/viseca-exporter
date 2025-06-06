package csv_test

import (
	"testing"

	"github.com/anothertobi/viseca-exporter/internal/csv"
	"github.com/anothertobi/viseca-exporter/pkg/viseca"
	"github.com/stretchr/testify/assert"
)

var inputTransaction = viseca.Transaction{
	TransactionID:    "AUTH8c919db2-1c23-43f1-8862-61c31336d9b6",
	CardID:           "0000000AAAAA0000",
	MaskedCardNumber: "XXXXXXXXXXXX0000",
	CardName:         "Mastercard",
	Date:             "2021-10-20T17:05:44",
	ShowTimestamp:    true,
	Amount:           50.55,
	Currency:         "CHF",
	OriginalAmount:   50.55,
	OriginalCurrency: "CHF",
	MerchantName:     "Aldi Suisse 00",
	PrettyName:       "ALDI",
	MerchantPlace:    "",
	IsOnline:         false,
	PFMCategory: viseca.PFMCategory{
		ID:                  "cv_groceries",
		Name:                "Lebensmittel",
		LightColor:          "#E2FDD3",
		MediumColor:         "#A5D58B",
		Color:               "#51A127",
		ImageURL:            "https://api.one.viseca.ch/v1/media/categories/icon_with_background/ic_cat_tile_groceries_v2.png",
		TransparentImageURL: "https://api.one.viseca.ch/v1/media/categories/icon_without_background/ic_cat_tile_groceries_v2.png",
	},
	StateType: "authorized",
	Details:   "Aldi Suisse 00",
	Type:      "merchant",
	IsBilled:  false,
	Links: viseca.TransactionLinks{
		Transactiondetails: "/v1/card/0000000AAAAA0000/transaction/AUTH8c919db2-1c23-43f1-8862-61c31336d9b6",
	},
}

func TestTransactionString(t *testing.T) {
	expected := `"AUTH8c919db2-1c23-43f1-8862-61c31336d9b6","2021-10-20T17:05:44","ALDI","50.55","cv_groceries","Lebensmittel"`

	assert.Equal(t, expected, csv.TransactionString(inputTransaction))
}

func TestTransactionsString(t *testing.T) {
	inputTransactions := []viseca.Transaction{inputTransaction}
	expected :=
		`"TransactionID","Date","Merchant","Amount","PFMCategoryID","PFMCategoryName"` +
			"\n" +
			`"AUTH8c919db2-1c23-43f1-8862-61c31336d9b6","2021-10-20T17:05:44","ALDI","50.55","cv_groceries","Lebensmittel"` +
			"\n"

	assert.Equal(t, expected, csv.TransactionsString(inputTransactions))
}

func TestForeignCurrencyTransaction(t *testing.T) {
	foreignTransaction := viseca.Transaction{
		TransactionID:    "TRX2025051200004466612",
		Date:             "2025-05-12T09:01:20+02:00",
		MerchantName:     "CLAUDE.AI SUBSCRIPTION",
		PrettyName:       "Claude.ai",
		IsOnline:         true,
		Amount:           17.15,
		Currency:         "CHF",
		OriginalAmount:   20.00,
		OriginalCurrency: "USD",
		PFMCategory: viseca.PFMCategory{
			ID:   "entertainment_and_leisure",
			Name: "Entertainment & Leisure",
		},
		StateType: "booked",
		Details:   "CLAUDE.AI SUBSCRIPTION",
		Type:      "merchant",
		IsBilled:  true,
	}

	expected := `"TRX2025051200004466612","2025-05-12T09:01:20+02:00","Claude.ai","17.15","CHF","20.00","USD","entertainment_and_leisure","Entertainment & Leisure"`

	options := csv.TransactionOptions{
		IncludeForeignCurrency: true,
	}

	assert.Equal(t, expected, csv.TransactionStringWithOptions(foreignTransaction, options))
}
