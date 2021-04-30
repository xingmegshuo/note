package config

import (
	"github.com/tal-tech/go-zero/core/stores/cache"
	"github.com/tal-tech/go-zero/rest"
)

type Config struct {
	rest.RestConf
	DataSource string
	CacheRedis cache.CacheConf
}
