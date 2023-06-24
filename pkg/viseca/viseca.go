package viseca

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

const (
	defaultBaseURL = "https://api.one.viseca.ch/v1/"
)

// Client represents an API client.
type Client struct {
	HTTPClient *http.Client
	BaseURL    *url.URL
}

// NewClient returns a new Viseca API client.
func NewClient(httpClient *http.Client) *Client {
	if httpClient == nil {
		httpClient = &http.Client{}
	}
	baseURL, _ := url.Parse(defaultBaseURL)

	client := &Client{HTTPClient: httpClient, BaseURL: baseURL}

	return client
}

const listOptionsDateFormat = "2006-01-02T15:04:05Z"

// ListOptions holds the options for list actions.
type ListOptions struct {
	Offset    int
	PageSize  int
	StateType string
	DateTo    time.Time
	DateFrom  time.Time
}

// NewDefaultListOptions creates new default ListOptions.
func NewDefaultListOptions() ListOptions {
	listOptions := ListOptions{}
	listOptions.Offset = 0
	listOptions.PageSize = 100
	listOptions.StateType = "unknown"

	return listOptions
}

func addListOptions(url *url.URL, listOptions ListOptions) {
	query := url.Query()
	query.Add("offset", strconv.Itoa(listOptions.Offset))
	query.Add("pagesize", strconv.Itoa(listOptions.PageSize))
	query.Add("statetype", listOptions.StateType)
	if !listOptions.DateFrom.IsZero() {
		query.Add("dateFrom", listOptions.DateFrom.Format(listOptionsDateFormat))
	}
	if !listOptions.DateTo.IsZero() {
		query.Add("dateTo", listOptions.DateTo.Format(listOptionsDateFormat))
	}

	url.RawQuery = query.Encode()
}

// CardListOptions known CreditIndicators are "credit" and "debit".
type CardListOptions struct {
	CreditIndicators []string
}

// NewDefaultCardListOptions creates new default CardListOptions.
func NewDefaultCardListOptions() CardListOptions {
	cardListOptions := CardListOptions{}
	cardListOptions.CreditIndicators = []string{"credit", "debit"}

	return cardListOptions
}

func addCardListOptions(url *url.URL, cardListOptions CardListOptions) {
	query := url.Query()

	for _, creditIndicator := range cardListOptions.CreditIndicators {
		query.Add("creditIndicators", creditIndicator)
	}

	url.RawQuery = query.Encode()
}

// Do returns the HTTP response and decodes the JSON response body into responseBody.
func (client *Client) Do(ctx context.Context, request *http.Request, responseBody interface{}) (*http.Response, error) {
	request = request.WithContext(ctx)
	response, err := client.HTTPClient.Do(request)
	if err != nil {
		return response, err
	}
	defer response.Body.Close()

	if response.StatusCode >= 200 && response.StatusCode < 300 {
		if responseBody != nil {
			err = json.NewDecoder(response.Body).Decode(responseBody)
		}
	} else {
		err = fmt.Errorf("api call failed with status: %s", response.Status)
	}

	return response, err
}

// NewRequest returns a new request with the given path (adds the base URL) and method. It further encodes the requestBody as JSON.
func (client *Client) NewRequest(path string, method string, requestBody interface{}) (*http.Request, error) {
	requestURL := client.BaseURL.JoinPath(path)

	var buffer io.ReadWriter

	if requestBody != nil {
		buffer := &bytes.Buffer{}
		err := json.NewEncoder(buffer).Encode(requestBody)
		if err != nil {
			return nil, err
		}
	}

	request, err := http.NewRequest(method, requestURL.String(), buffer)
	if err != nil {
		return nil, err
	}

	request.Header.Add("Accept", "application/json")

	return request, err
}
