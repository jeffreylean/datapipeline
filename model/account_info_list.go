package model

import (
	"strconv"
	"time"

	xmlType "data-pipeline/types"

	"github.com/si3nloong/sqlike/types"
)

type AccountInfoList struct {
	Key               *types.Key `json:"key" sqlike:"$Key"`
	ReferenceDate     *time.Time `json:"referenceDate"`
	DebitAccount      string     `json:"debitAccount"`
	CreditAccount     string     `json:"creditAccount"`
	LocalCurrency     string     `json:"localCurrency"`
	LocalDebitAmount  float64    `json:"localDebitAmount"`
	LocalCreditAmount float64    `json:"localCreditAmount"`
	HostCurrency      string     `json:"hostCurrency"`
	HostDebitAmount   float64    `json:"hostDebitAmount"`
	HostCreditAmount  float64    `json:"hostCreditAmount"`
	ExchangeRate      float64    `json:"exchangeRate"`
	EffectiveDate     *time.Time `json:"effectiveDate"`
	CreatedDateTime   time.Time  `json:"createdAt"`
}

func (a *AccountInfoList) Transform(i xmlType.AccountInfoList, transactionKey *types.Key) {
	a.Key = types.NewIDKey("AccountInfoList", transactionKey)
	if i.ReferenceDate != "" {
		referenceDate, _ := time.Parse("2006-01-02", i.ReferenceDate)
		a.ReferenceDate = &referenceDate
	}
	if i.LocalDebitAmount != "" {
		amount, _ := strconv.ParseFloat(i.LocalDebitAmount, 64)
		a.LocalDebitAmount = amount
	}
	if i.LocalCreditAmount != "" {
		amount, _ := strconv.ParseFloat(i.LocalCreditAmount, 64)
		a.LocalCreditAmount = amount
	}
	a.HostCurrency = i.HostCurrency
	if i.HostDebitAmount != "" {
		amount, _ := strconv.ParseFloat(i.HostDebitAmount, 64)
		a.HostDebitAmount = amount
	}
	if i.HostCreditAmount != "" {
		amount, _ := strconv.ParseFloat(i.HostCreditAmount, 64)
		a.HostCreditAmount = amount
	}
	if i.ExchangeRate != "" {
		amount, _ := strconv.ParseFloat(i.ExchangeRate, 64)
		a.ExchangeRate = amount
	}
	if i.EffectiveDate != "" {
		effectiveDate, _ := time.Parse("2006-01-02", i.EffectiveDate)
		a.EffectiveDate = &effectiveDate
	}
	a.DebitAccount = i.DebitAccount
	a.CreditAccount = i.CreditAccount
	a.LocalCurrency = i.LocalCurrency
	a.CreatedDateTime = time.Now().UTC()
}
