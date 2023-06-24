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

type User struct {
	VisecaOneID         string `json:"visecaOneId"`
	Email               string `json:"email"`
	EmailStatus         string `json:"emailStatus"`
	Language            string `json:"language"`
	FirstName           string `json:"firstName"`
	LastName            string `json:"lastName"`
	Gender              string `json:"gender"`
	MaskedPhoneNumber   string `json:"maskedPhoneNumber"`
	LastLoginDate       string `json:"lastLoginDate"`
	LastLogoutType      string `json:"lastLogoutType"`
	DefaultChannelType  string `json:"defaultChannelType"`
	AppDescription      string `json:"appDescription"`
	AppRegistrationDate string `json:"appRegistrationDate"`
}

type Card struct {
	ID                          string           `json:"cardId"`
	CardSwitch                  CardSwitch       `json:"switch"`
	MaskedCardNumber            string           `json:"maskedCardNumber"`
	CardAccountNr               string           `json:"cardAccountNr"`
	CardHolder                  CardHolder       `json:"cardHolder"`
	CardStatus                  CardStatus       `json:"cardStatus"`
	CardName                    string           `json:"cardName"`
	CardDescription             string           `json:"cardDescription"`
	ExpirationDate              string           `json:"expirationDate"`
	ProductType                 string           `json:"productType"`
	ProductLine                 string           `json:"productLine"`
	CreditIndicator             string           `json:"creditIndicator"`
	AvailableReplacementReasons []string         `json:"availableReplacementReasons"`
	BonusProgram                []string         `json:"bonusProgram"`
	MainBonusProgram            string           `json:"mainBonusProgram"`
	Currency                    string           `json:"currency"`
	CardLimit                   float32          `json:"cardLimit"`
	ActiveCurrency              string           `json:"activeCurrency"`
	ActiveLimit                 float32          `json:"activeLimit"`
	IsSelfIssued                bool             `json:"isSelfIssued"`
	CardScheme                  string           `json:"cardScheme"`
	EmbossingLine               string           `json:"embossingLine"`
	CardType                    CardType         `json:"cardType"`
	CallCenter                  string           `json:"callCenter"`
	CardImageDetails            CardImageDetails `json:"cardImageDetails"`
	CardLinks                   CardLinks        `json:"links"`
	CardGrants                  CardGrants       `json:"grants"`
}

type CardSwitch struct {
	Reason string `json:"reason"`
}

type CardHolder struct {
	Firstname     string `json:"firstname"`
	Lastname      string `json:"lastname"`
	BirthDate     string `json:"birthDate"`
	Nationality   string `json:"nationality"`
	IsCurrentUser bool   `json:"isCurrentUser"`
}

type CardStatus struct {
	Value       string `json:"value"`
	AdvValue    string `json:"advValue"`
	Description string `json:"description"`
	ChangeDate  string `json:"changeDate"`
}

type CardType struct {
	Value       string `json:"value"`
	Description string `json:"description"`
}

type CardImageDetails struct {
	URL                     string `json:"url"`
	TemplateName            string `json:"templateName"`
	Category                string `json:"category"`
	Status                  string `json:"status"`
	DenialReason            string `json:"denialReason"`
	ReplacementAvailability string `json:"replacementAvailability"`
	UploadContext           string `json:"uploadContext"`
	LastStatusUpdate        string `json:"lastStatusUpdate"`
}

type CardLinks struct {
	CardDetails           string `json:"carddetails"`
	CardImage             string `json:"cardimage"`
	CardSwitcherLogoImage string `json:"cardswitcherlogoimage"`
	CockpitLogoImage      string `json:"cockpitlogoimage"`
}

type CardGrants struct {
	CanSurprizeRead                  bool `json:"canSurprizeRead"`
	CanAccountDetailsRead            bool `json:"canAccountDetailsRead"`
	CanStatementSettingsRead         bool `json:"canStatementSettingsRead"`
	CanStatementSettingsUpdate       bool `json:"canStatementSettingsUpdate"`
	CanStatementDetailRead           bool `json:"canStatementDetailRead"`
	CanTransactionNotificationRead   bool `json:"canTransactionNotificationRead"`
	CanTransactionNotificationUpdate bool `json:"canTransactionNotificationUpdate"`
	CanMasterpassRead                bool `json:"canMasterpassRead"`
	CanMasterpassUpdate              bool `json:"canMasterpassUpdate"`
	CanCardFreeze                    bool `json:"canCardFreeze"`
	CanCardPINRequest                bool `json:"canCardPINRequest"`
	CanSmsSettingsRead               bool `json:"canSmsSettingsRead"`
	CanSmsSettingsUpdate             bool `json:"canSmsSettingsUpdate"`
	CanReplaceCard                   bool `json:"canReplaceCard"`
	CanCardControlsRead              bool `json:"canCardControlsRead"`
	CanCardControlsUpdate            bool `json:"canCardControlsUpdate"`
	CanCouponsRead                   bool `json:"canCouponsRead"`
	CanPanCvvPinRead                 bool `json:"canPanCvvPinRead"`
}
