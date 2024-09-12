package server

import (
	"github.com/go-bamboo/pkg/rpc"
)

// NewGRPCServer new a gRPC server.
func NewGRPCServer() *rpc.Server {
	srv := rpc.NewServer(&rpc.Conf{
		Address: ":9000",
	})
	return srv
}
