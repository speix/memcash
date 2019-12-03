package services

import (
	"github.com/speix/memcash/app"
	"github.com/speix/memcash/app/config"
)

type Engine struct {
	*config.GRPCNetwork
	*app.Cache
}
