package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type Transactions struct {
	Transactions []Transaction `json:"list"`
}

type Transaction struct {
	Date         string  `json:"date"`
	Amount       float64 `json:"amount"`
	MerchantName string  `json:"merchantName"`
}

const URL_PRE = "https://api.one-digitalservice.ch/v1/card/"
const URL_POST = "/transactions?stateType=unknown&offset=0&pagesize=1000"

// arg0: cardID
// arg1: sessionCookie (e.g. `AL_SESS-S=...`)
func main() {
	transactions := getTransactions(os.Args[0], os.Args[1])
	printTransactions(transactions)
}

func getTransactions(cardID, sessionCookie string) Transactions {
	var transactions Transactions

	client := &http.Client{}

	req, err := http.NewRequest("GET", URL_PRE+cardID+URL_POST, nil)
	check(err)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Cookie", sessionCookie)

	resp, err := client.Do(req)

	check(err)
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	check(err)

	err = json.Unmarshal(data, &transactions)
	check(err)

	return transactions
}

func printTransactions(transactions Transactions) {
	fmt.Println("\"Date\",\"Payee\",\"Amount\"")

	for _, v := range transactions.Transactions {
		fmt.Printf("\"%s\",\"%s\",\"%f\"\n", v.Date, v.MerchantName, v.Amount)
	}
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
