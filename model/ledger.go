package model

import (
	"time"

	xmlType "data-pipeline/types"

	"github.com/si3nloong/sqlike/types"
)

type Ledger struct {
	Key             *types.Key `json:"key" sqlike:"$Key"`
	CompanyCode     string     `json:"companyCode"`
	AccountingDate  string     `json:"accountingDate"`
	CreatedDateTime time.Time  `json:"createdAt"`
}

func (l *Ledger) Transform(i xmlType.Header) {
	l.Key = types.NewIDKey("Ledger", nil)
	l.CompanyCode = i.CompanyCode
	l.AccountingDate = i.AccountingDate
	l.CreatedDateTime = time.Now().UTC()
}
