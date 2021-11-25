package repository

import (
	"data-pipeline/repository/function"

	"github.com/si3nloong/sqlike/sqlike"
)

type Repository struct {
	MySQL           *sqlike.Database
	Client          *sqlike.Client
	AccountInfoList AccountInfoList
	Ledger          Ledger
	Transaction     Transaction
}

func NewMySQL(conn *sqlike.Client, db *sqlike.Database) Repository {
	return Repository{db, conn, function.AccountInfoList{MySQL: db}, function.Ledger{MySQL: db}, function.Transaction{MySQL: db}}
}
