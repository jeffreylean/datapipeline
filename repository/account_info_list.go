package repository

import (
	"data-pipeline/model"
)

type AccountInfoList interface {
	Migration() error
	CreateAccountInfoList(accountInfo *model.AccountInfoList) error
}
