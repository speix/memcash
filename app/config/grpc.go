package config

import (
	"log"
	"net"
	"os"
)

type GRPCNetwork struct {
	net.Listener
}

type GRPCConfig struct {
	Network string
	Port    string
}

func GetGRPCConfig() *GRPCConfig {
	return &GRPCConfig{
		Network: "tcp",
		Port:    os.Getenv("MEMCASH_GRPC_PORT"),
	}
}

func NewGRPCListener(config *GRPCConfig) *GRPCNetwork {

	listener, err := net.Listen(config.Network, ":"+config.Port)
	if err != nil {
		log.Fatal("Failed to listen", err)
	}

	return &GRPCNetwork{Listener: listener}
}

func GetGRPCUser() string {
	return os.Getenv("MEMCASH_GRPC_USER")
}

func GetGRPCPass() string {
	return os.Getenv("MEMCASH_GRPC_PASS")
}
