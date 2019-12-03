package middlwares

import (
	"context"
	"encoding/base64"
	"strings"

	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	"github.com/speix/memcash/app/config"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Authorize authorizes a gRPC call to this service
func Authorize(ctx context.Context) (context.Context, error) {

	token, err := grpc_auth.AuthFromMD(ctx, "Basic")
	if err != nil {
		return nil, status.Error(codes.Unauthenticated, err.Error())
	}

	c, err := base64.StdEncoding.DecodeString(token)
	if err != nil {
		return nil, status.Error(codes.Unauthenticated, "Invalid base64 authorization encoding")
	}

	cs := string(c)
	s := strings.IndexByte(cs, ':')
	if s < 0 {
		return nil, status.Error(codes.Unauthenticated, "Invalid base64 authorization format")
	}

	user, pass := cs[:s], cs[s+1:]
	if user != config.GetGRPCUser() || pass != config.GetGRPCPass() {
		return nil, status.Error(codes.Unauthenticated, "Invalid auth micro-service access")
	}

	return ctx, nil
}
