package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"net/http/cookiejar"
	"os"
	"strings"

	"github.com/anothertobi/viseca-exporter/internal/csv"
	"github.com/anothertobi/viseca-exporter/pkg/viseca"
)

const sessionCookieName = "AL_SESS-S"

// arg0: cardID
// arg1: sessionCookie (example: `AL_SESS-S=...`)
// arg2: includeForeignCurrency (example: "true")
func main() {
	if len(os.Args) < 3 {
		log.Fatal("card ID and session cookie args required")
	}
	visecaClient, err := initClient(os.Args[2])
	if err != nil {
		log.Fatalf("error initializing Viseca API client: %v", err)
	}

	ctx := context.Background()

	transactions, err := visecaClient.ListAllTransactions(ctx, os.Args[1])
	if err != nil {
		log.Fatalf("error listing all transactions: %v", err)
	}

	includeForeignCurrency := false
	if len(os.Args) == 4 && strings.ToLower(os.Args[3]) == "true" {
		includeForeignCurrency = true
	}
	options := csv.TransactionOptions{
		IncludeForeignCurrency: includeForeignCurrency,
	}

	fmt.Println(csv.TransactionsStringWithOptions(transactions, options))
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

func extractSessionCookieValue(sessionCookie string) string {
	return strings.TrimPrefix(sessionCookie, fmt.Sprintf("%s=", sessionCookieName))
}
