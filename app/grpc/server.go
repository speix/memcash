package grpc

import (
	"log"

	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"

	"github.com/speix/memcash/app/grpc/middlwares"
	v1 "github.com/speix/memcash/app/pb/v1"
	"github.com/speix/memcash/app/services"
	"google.golang.org/grpc"
)

type Server struct {
	MemcashService *services.Engine
}

func StartGRPCServer(engine *services.Engine) {

	s := grpc.NewServer(
		grpc.UnaryInterceptor(
			grpc_auth.UnaryServerInterceptor(middlwares.Authorize),
		),
	)

	v1.RegisterMemcashServiceServer(s, &Server{MemcashService: engine})
	v1.RegisterHealthServiceServer(s, NewHealthService(s))

	if err := s.Serve(engine.GRPCNetwork.Listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
