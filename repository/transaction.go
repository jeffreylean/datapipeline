package repository

import (
	"data-pipeline/model"
)

type Transaction interface {
	Migration() error
	CreateTransaction(transaction *model.Transaction) error
}
