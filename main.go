package main

import (
	"data-pipeline/bootstrap"
	"data-pipeline/middleware"

	handler "data-pipeline/handler/v1"

	"github.com/labstack/echo/v4"
)

func main() {
	bs := bootstrap.New()
	defer bs.Close()

	h := handler.New(bs)

	e := echo.New()

	v1 := e.Group("/v1")

	//migration
	migration := e.Group("/migration")
	migration.POST("/account-info-list", h.MigrateAccountInfoList)
	migration.POST("/transaction", h.MigrateTransaction)
	migration.POST("/ledger", h.MigrateLedger)

	v1.POST("/transfer", h.DatalakeMigration, middleware.ErrorNotify())

	e.Logger.Fatal(e.Start(":1323"))
}
