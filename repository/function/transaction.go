package function

import (
	"context"
	"data-pipeline/model"

	"github.com/si3nloong/sqlike/sqlike"
)

type Transaction struct {
	MySQL *sqlike.Database
}

//migration
func (r Transaction) Migration() error {
	ctx := context.Background()
	table := r.MySQL.Table("Transaction")
	table.Migrate(ctx, model.Transaction{})
	return nil
}

//Create Transaction
func (r Transaction) CreateTransaction(transaction *model.Transaction) error {
	table := r.MySQL.Table("Transaction")
	_, err := table.InsertOne(context.Background(), transaction)
	if err != nil {
		return err
	}
	return nil
}
