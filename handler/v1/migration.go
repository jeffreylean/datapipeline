package v1

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h Handler) MigrateAccountInfoList(c echo.Context) error {
	if err := h.repository.AccountInfoList.Migration(); err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, nil)
}

func (h Handler) MigrateTransaction(c echo.Context) error {
	if err := h.repository.Transaction.Migration(); err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, nil)
}

func (h Handler) MigrateLedger(c echo.Context) error {
	if err := h.repository.Ledger.Migration(); err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, nil)
}
