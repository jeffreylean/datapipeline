package v1

import (
	"data-pipeline/model"
	"data-pipeline/response"
	"data-pipeline/service"
	xmlType "data-pipeline/types"
	"encoding/json"
	"encoding/xml"
	"net/http"
	"time"

	"data-pipeline/response/errcode"

	"github.com/ivpusic/grpool"
	"github.com/labstack/echo/v4"
)

func (h Handler) DatalakeMigration(c echo.Context) error {
	dates := make([]string, 0)
	//check any failed previous job
	d, err := h.redis.Get("FAILED_JOB").Result()
	if err != nil {
		b, _ := json.Marshal(dates)
		h.redis.Set("FAILED_JOB", b, time.Hour*48)
		d, _ := h.redis.Get("FAILED_JOB").Result()
		json.Unmarshal([]byte(d), &dates)
	} else {
		json.Unmarshal([]byte(d), &dates)
	}
	//add today dates
	location, _ := time.LoadLocation("Asia/Kuala_Lumpur")
	todayDate := time.Now().In(location).Format("2006-01-02")
	dates = append(dates, todayDate)

	temp := dates
	for x := 0; x < len(dates); x++ {
		data, err := service.XMLRetriever(dates[x])
		if err != nil {
			temp = append(temp, todayDate)
			b, _ := json.Marshal(temp)
			h.redis.Set("FAILED_JOB", b, time.Hour*48)
			return c.JSON(http.StatusInternalServerError, response.Exception{Code: errcode.XMLExtractFailed, Message: err.Error()})
		}

		record := &xmlType.Ledger{}
		if err := xml.Unmarshal(data, record); err != nil {
			temp = append(temp, todayDate)
			b, _ := json.Marshal(temp)
			h.redis.Set("FAILED_JOB", b, time.Hour*48)
			return c.JSON(http.StatusInternalServerError, response.Exception{Code: errcode.UnmarshalFailed, Message: err.Error()})
		}
		//migrate ledger
		ledger := new(model.Ledger)
		ledger.Transform(record.Header)
		h.repository.Ledger.CreateLedger(ledger)

		// migrate to account info list
		transactions := record.Transactions.Transaction
		pool := grpool.NewPool(20, 20)
		defer pool.Release()
		pool.WaitCount(len(transactions))

		for i := range transactions {
			pool.JobQueue <- func(index int) func() {
				return func() {
					defer pool.JobDone()
					transaction := new(model.Transaction)
					transaction.Transform(transactions[i], ledger.Key)
					h.repository.Transaction.CreateTransaction(transaction)
					for _, each := range transactions[i].AccountInfoLists.AccountInfoList {
						accountInfo := new(model.AccountInfoList)
						accountInfo.Transform(each, transaction.Key)
						h.repository.AccountInfoList.CreateAccountInfoList(accountInfo)
					}
				}
			}(i)
		}

		pool.WaitAll()
		dates = append(dates[:x], dates[x+1:]...)
		temp = dates
		x--
	}

	b, _ := json.Marshal(temp)
	h.redis.Set("FAILED_JOB", b, time.Hour*48)

	return c.JSON(http.StatusOK, nil)
}
