package main

import (
	"github.com/speix/memcash/app/config"
	"github.com/speix/memcash/app/grpc"
	"github.com/speix/memcash/app/services"
)

func main() {

	grpcNetwork := config.NewGRPCListener(config.GetGRPCConfig())
	cache := services.NewCache()

	defer grpcNetwork.Listener.Close()

	memcash := &services.Engine{
		GRPCNetwork: grpcNetwork,
		Cache:       cache,
	}

	grpc.StartGRPCServer(memcash)
	memcash.Supervise()
}
