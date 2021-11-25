package model

import (
	"time"

	xmlType "data-pipeline/types"

	"github.com/si3nloong/sqlike/types"
)

type Transaction struct {
	Key                  *types.Key `json:"key" sqlike:"$Key"`
	JournalCode          string     `json:"journalCode"`
	JournalDescription   string     `json:"journalDescription"`
	AccountEvent         string     `json:"accountEvent"`
	AccountType          string     `json:"accountType"`
	DepartureStation     string     `json:"departureStation"`
	RouteAndFlightNumber string     `json:"routeAndFlightNumber"`
	TailNumber           string     `json:"tailNumber"`
	ReferenceCode        string     `json:"referenceCode"`
	CreatedDateTime      time.Time  `json:"createdAt"`
}

func (t *Transaction) Transform(i xmlType.Transaction, ledgerKey *types.Key) {
	t.Key = types.NewIDKey("Transaction", ledgerKey)
	t.JournalCode = i.JournalCode
	t.JournalDescription = i.JournalDescription
	t.AccountEvent = i.AccountEvent
	t.AccountType = i.AccountType
	t.DepartureStation = i.DepartureStation
	t.RouteAndFlightNumber = i.RouteAndFlightNumber
	t.TailNumber = i.TailNumber
	t.ReferenceCode = i.ReferenceCode
	t.CreatedDateTime = time.Now().UTC()
}
