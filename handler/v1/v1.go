package v1

import (
	"data-pipeline/bootstrap"
	"data-pipeline/repository"

	"github.com/go-redis/redis"
)

// Handler :
type Handler struct {
	repository repository.Repository
	redis      redis.Client
}

// New :
func New(bs *bootstrap.Bootstrap) Handler {
	return Handler{
		repository: bs.Repository,
		redis:      *bs.Redis,
	}
}
