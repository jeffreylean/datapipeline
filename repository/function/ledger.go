package function

import (
	"context"
	"data-pipeline/model"

	"github.com/si3nloong/sqlike/sqlike"
)

type Ledger struct {
	MySQL *sqlike.Database
}

//migration
func (r Ledger) Migration() error {
	ctx := context.Background()
	table := r.MySQL.Table("Ledger")
	table.Migrate(ctx, model.Ledger{})
	return nil
}

//Create Ledger
func (r Ledger) CreateLedger(ledger *model.Ledger) error {
	table := r.MySQL.Table("Ledger")
	_, err := table.InsertOne(context.Background(), ledger)
	if err != nil {
		return err
	}
	return nil
}
