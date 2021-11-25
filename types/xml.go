package types

import "encoding/xml"

type Ledger struct {
	XMLName      xml.Name     `xml:"Ledger"`
	Header       Header       `xml:"Header"`
	Transactions Transactions `xml:"Transactions"`
}

type Header struct {
	CompanyCode    string `xml:"CompanyCode"`
	AccountingDate string `xml:"AccountingDate"`
}

type Transaction struct {
	JournalCode          string           `xml:"JournalCode"`
	JournalDescription   string           `xml:"JournalDescription"`
	AccountEvent         string           `xml:"AccountEvent"`
	AccountType          string           `xml:"AccountType"`
	DepartureStation     string           `xml:"DepartureStation"`
	RouteAndFlightNumber string           `xml:"RouteAndFlightNumber"`
	TailNumber           string           `xml:"TailNumber"`
	ReferenceCode        string           `xml:"ReferenceCode"`
	AccountInfoLists     AccountInfoLists `xml:"AccountInfoLists"`
}

type Transactions struct {
	Transaction []Transaction `xml:"Transaction"`
}

type AccountInfoLists struct {
	AccountInfoList []AccountInfoList `xml:"AccountInfoList"`
}

type AccountInfoList struct {
	ReferenceDate     string `xml:"ReferenceDate"`
	DebitAccount      string `xml:"DebitAccount"`
	CreditAccount     string `xml:"CreditAccount"`
	LocalCurrency     string `xml:"LocalCurrency"`
	LocalDebitAmount  string `xml:"LocalDebitAmount"`
	LocalCreditAmount string `xml:"LocalCreditAmount"`
	HostCurrency      string `xml:"HostCurrency"`
	HostDebitAmount   string `xml:"HostDebitAmount"`
	HostCreditAmount  string `xml:"HostCreditAmount"`
	ExchangeRate      string `xml:"ExchangeRate"`
	EffectiveDate     string `xml:"EffectiveDate"`
}
