package bootstrap

import (
	"context"
	"data-pipeline/config"
	"data-pipeline/repository"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/si3nloong/sqlike/sqlike"
	"github.com/si3nloong/sqlike/sqlike/options"
)

func (bs *Bootstrap) initSQL() *Bootstrap {
	client := sqlike.MustConnect(context.Background(), "mysql",
		options.Connect().
			SetHost(config.DBHost).
			SetPort(config.DBPort).
			SetUsername(config.DBUserName).
			SetPassword(config.DBPassword),
	)
	bs.MySQL = client
	bs.Repository = repository.NewMySQL(bs.MySQL, bs.MySQL.Database(config.DBName))

	os.Stdout.WriteString("Connected to MySQL!\n")

	return bs
}
