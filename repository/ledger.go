package repository

import (
	"data-pipeline/model"
)

type Ledger interface {
	Migration() error
	CreateLedger(ledger *model.Ledger) error
}
