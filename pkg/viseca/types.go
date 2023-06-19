package viseca

type Transactions struct {
	TotalCount   int           `json:"totalCount"`
	Transactions []Transaction `json:"list"`
}

type Transaction struct {
	TransactionID    string           `json:"transactionId"`
	CardID           string           `json:"cardId"`
	MaskedCardNumber string           `json:"maskedCardNumber"`
	CardName         string           `json:"cardName"`
	Date             string           `json:"date"`
	ShowTimestamp    bool             `json:"showTimestamp"`
	Amount           float64          `json:"amount"`
	Currency         string           `json:"currency"`
	OriginalAmount   float64          `json:"originalAmount"`
	OriginalCurrency string           `json:"originalCurrency"`
	MerchantName     string           `json:"merchantName"`
	PrettyName       string           `json:"prettyName"`
	MerchantPlace    string           `json:"merchantPlace"`
	IsOnline         bool             `json:"isOnline"`
	PFMCategory      PFMCategory      `json:"pfmCategory"`
	StateType        string           `json:"stateType"`
	Details          string           `json:"details"`
	Type             string           `json:"type"`
	IsBilled         bool             `json:"isBilled"`
	Links            TransactionLinks `json:"links"`
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

type TransactionLinks struct {
	Transactiondetails string `json:"transactiondetails"`
}
