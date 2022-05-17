package models

type Currencies string

const (
	Tl   Currencies = "TRY"
	Usd             = "USD"
	Euro            = "EUR"
)

type AccountTypes string

const (
	Ind  AccountTypes = "individual"
	Corp              = "corporate"
)

type Account struct {
	AccountNumber int          `json:accountNumber`
	CurrencyCode  Currencies   `json:currencyCode`
	OwnerName     string       `json:ownerName`
	AccountType   AccountTypes `json:accountType`
	Balance       float32      `json:balance`
}
