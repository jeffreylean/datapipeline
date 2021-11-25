package bootstrap

import (
	"data-pipeline/repository"

	"github.com/go-redis/redis"
	"github.com/si3nloong/sqlike/sqlike"
)

type Bootstrap struct {
	Repository repository.Repository
	MySQL      *sqlike.Client
	Redis      *redis.Client
}

// New :
func New() *Bootstrap {

	bs := new(Bootstrap)
	bs.initSQL()
	bs.initRedis()

	return bs
}

// Close :
func (bs *Bootstrap) Close() error {

	if err := bs.Repository.Client.Close(); err != nil {
		return err
	}

	return nil
}
