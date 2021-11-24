package adapters

import (
	"context"
	"net"
	"sync"

	"google.golang.org/grpc"
)

var grpcOnce = &sync.Once{}
var grpcSingleton IAdapter

// GetGRPC provides the gRPC adapter singleton.
func GetGRPC() IAdapter {
	grpcOnce.Do(func() {
		grpcSingleton = &implGRPC{}
		grpcSingleton.init()
	})

	return grpcSingleton
}

// implGRPC implements IAdapter as a gRPC server.
type implGRPC struct {
	server *grpc.Server
}

func (i *implGRPC) Name() string {
	return "gRPC"
}

func (i *implGRPC) Start(ctx context.Context) error {
	if !conf.GRPCServer.Enabled {
		log.Infof("gRPC server is not enabled.")
		return nil
	}

	listener, err := net.Listen("tcp", conf.GRPCServer.Addr)
	if err != nil {
		log.Errorf("Failed to create the listener: %s", err.Error())
		return err
	}

	i.server = grpc.NewServer()
	/* Attach the services here. */

	log.Infof("Starting %s@%s gRPC server at %s", conf.Application.Name, conf.Application.Version, conf.GRPCServer.Addr)
	return i.server.Serve(listener)
}

func (i *implGRPC) Stop(ctx context.Context) error {
	if i.server == nil {
		log.Infof("Not doing anything because server pointer is nil.")
		return nil
	}

	log.Infof("Stopping the gRPC server...")
	i.server.GracefulStop()
	log.Infof("gRPC server stopped gracefully.")

	return nil
}

func (i *implGRPC) init() {}
