package main

import (
	"github.com/speix/memcash/app/config"
	"github.com/speix/memcash/app/grpc"
	"github.com/speix/memcash/app/services"
)

// Memory storage gRPC service with time persistence.
// Supports authentication using environment variables
// and health checking functionality in case it integrates
// with a container-orchestration system like Kubernetes.
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
