package function

import (
	"context"
	"data-pipeline/model"

	"github.com/si3nloong/sqlike/sqlike"
)

type AccountInfoList struct {
	MySQL *sqlike.Database
}

//migration
func (r AccountInfoList) Migration() error {
	ctx := context.Background()
	table := r.MySQL.Table("AccountInfoList")
	table.Migrate(ctx, model.AccountInfoList{})
	return nil
}

//Create AccountInfoList
func (r AccountInfoList) CreateAccountInfoList(accountInfo *model.AccountInfoList) error {
	table := r.MySQL.Table("AccountInfoList")
	_, err := table.InsertOne(context.Background(), accountInfo)
	if err != nil {
		return err
	}
	return nil
}
