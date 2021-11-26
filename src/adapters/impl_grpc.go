package adapters

import (
	"context"
	"net"

	"google.golang.org/grpc"
)

// implGRPC implements IAdapter as a gRPC server.
type implGRPC struct {
	server *grpc.Server
}

func (i *implGRPC) Name() string {
	return "gRPC"
}

func (i *implGRPC) Start(ctx context.Context) error {
	if !conf.GRPCServer.Enabled {
		log.Infof(ctx, "gRPC server is not enabled.")
		return nil
	}

	listener, err := net.Listen("tcp", conf.GRPCServer.Addr)
	if err != nil {
		log.Errorf(ctx, "Failed to create the listener: %s", err.Error())
		return err
	}

	i.server = grpc.NewServer()
	/* Attach the services here. */

	log.Infof(ctx, "Starting %s@%s gRPC server at %s", conf.Application.Name, conf.Application.Version, conf.GRPCServer.Addr)
	return i.server.Serve(listener)
}

func (i *implGRPC) Stop(ctx context.Context) error {
	if i.server == nil {
		log.Infof(ctx, "Not doing anything because server pointer is nil.")
		return nil
	}

	log.Infof(ctx, "Stopping the gRPC server...")
	i.server.GracefulStop()
	log.Infof(ctx, "gRPC server stopped gracefully.")

	return nil
}

func (i *implGRPC) init() {}
