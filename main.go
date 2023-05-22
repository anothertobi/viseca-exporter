package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"net/http/cookiejar"
	"os"
	"strings"

	"github.com/anothertobi/viseca-exporter/csv"
	"github.com/anothertobi/viseca-exporter/viseca"
)

const sessionCookieName = "AL_SESS-S"

// arg0: cardID
// arg1: sessionCookie (e.g. `AL_SESS-S=...`)
func main() {
	if len(os.Args) < 3 {
		log.Fatal("card ID and session cookie args required")
	}
	visecaClient, err := initClient(os.Args[2])
	if err != nil {
		log.Fatalf("error initializing Viseca API client: %v", err)
	}

	transactions, err := ListAllTransactions(visecaClient, os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(csv.TransactionsString(*transactions))
}

func initClient(sessionCookie string) (*viseca.Client, error) {
	cookieJar, err := cookiejar.New(nil)
	if err != nil {
		return nil, err
	}
	httpClient := http.Client{
		Jar: cookieJar,
	}
	visecaClient := viseca.NewClient(&httpClient)
	cookie := &http.Cookie{
		Name:  sessionCookieName,
		Value: extractSessionCookieValue(sessionCookie),
	}
	httpClient.Jar.SetCookies(visecaClient.BaseURL, []*http.Cookie{cookie})

	return visecaClient, nil
}

func ListAllTransactions(visecaClient *viseca.Client, cardID string) (*viseca.Transactions, error) {
	ctx := context.Background()
	listOptions := viseca.NewDefaultListOptions()
	transactions, err := visecaClient.ListTransactions(ctx, cardID, listOptions)
	if err != nil {
		return nil, err
	}
	for listOptions.Offset+listOptions.PageSize < transactions.TotalCount {
		listOptions.Offset += listOptions.PageSize
		transactionsPage, err := visecaClient.ListTransactions(ctx, cardID, listOptions)
		if err != nil {
			return nil, err
		}
		transactions.Transactions = append(transactions.Transactions, transactionsPage.Transactions...)
	}

	return transactions, nil
}

func extractSessionCookieValue(sessionCookie string) string {
	return strings.TrimPrefix(sessionCookie, fmt.Sprintf("%s=", sessionCookieName))
}
