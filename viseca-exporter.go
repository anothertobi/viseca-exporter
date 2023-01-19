package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

type Transactions struct {
	Transactions []Transaction `json:"list"`
}

type Transaction struct {
	TransactionID    string      `json:"transactionId"`
	CardID           string      `json:"cardId"`
	MaskedCardNumber string      `json:"maskedCardNumber"`
	CardName         string      `json:"cardName"`
	Date             string      `json:"date"`
	ShowTimestamp    bool        `json:"showTimestamp"`
	Amount           float64     `json:"amount"`
	Currency         string      `json:"currency"`
	OriginalAmount   float64     `json:"originalAmount"`
	OriginalCurrency string      `json:"originalCurrency"`
	MerchantName     string      `json:"merchantName"`
	PrettyName       string      `json:"prettyName"`
	MerchantPlace    string      `json:"merchantPlace"`
	IsOnline         bool        `json:"isOnline"`
	PFMCategory      PFMCategory `json:"pfmCategory"`
	StateType        string      `json:"stateType"`
	Details          string      `json:"details"`
	Type             string      `json:"type"`
	IsBilled         bool        `json:"isBilled"`
	Links            Links       `json:"links"`
}

type PFMCategory struct {
	ID                  string `json:"id"`
	Name                string `json:"name"`
	LightColor          string `json:"lightColor"`
	MediumColor         string `json:"mediumColor"`
	Color               string `json:"color"`
	ImageURL            string `json:"imageUrl"`
	TransparentImageURL string `json:"transparentImageUrl"`
}

type Links struct {
	Transactiondetails string `json:"transactiondetails"`
}

const URL_PRE = "https://api.one.viseca.ch/v1/card/"
const URL_POST = "/transactions?stateType=unknown&offset=0&pagesize=1000"

// arg0: cardID
// arg1: sessionCookie (e.g. `AL_SESS-S=...`)
func main() {
	if len(os.Args) < 3 {
		log.Fatal("card ID and session cookie args required")
	}
	transactions, err := getTransactions(os.Args[1], os.Args[2])
	if err != nil {
		log.Fatal(err)
	}
	printTransactions(transactions)
}

func getTransactions(cardID, sessionCookie string) (Transactions, error) {
	transactions := Transactions{}

	client := &http.Client{}

	url := URL_PRE + cardID + URL_POST

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return transactions, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Cookie", sessionCookie)

	resp, err := client.Do(req)
	if err != nil {
		return transactions, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return transactions, fmt.Errorf("request failed with status \"%s\"", resp.Status)
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return transactions, err
	}

	err = json.Unmarshal(data, &transactions)
	if err != nil {
		return transactions, err
	}

	return transactions, nil
}

func printTransactions(transactions Transactions) {
	fmt.Println("\"TransactionID\",\"Date\",\"Merchant\",\"Amount\",\"PFMCategoryID\",\"PFMCategoryName\"")

	for _, v := range transactions.Transactions {
		fmt.Printf("\"%s\",\"%s\",\"%s\",\"%f\",\"%s\",\"%s\"\n", v.TransactionID, v.Date, getPrettiestMerchantName(v), v.Amount, v.PFMCategory.ID, v.PFMCategory.Name)
	}
}

func getPrettiestMerchantName(transaction Transaction) string {
	if transaction.PrettyName != "" {
		return transaction.PrettyName
	} else {
		return transaction.MerchantName
	}
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
