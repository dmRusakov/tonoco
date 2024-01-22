package appInit

import (
	"google.golang.org/grpc"
	"net"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc/credentials/insecure"
)

// start grpc server
func (a *App) GrpcServerStart()
	// create grpc server
	grpcServer := grpc.NewServer()
	listen, err := net.Listen(
		"tcp",
		fmt.Sprintf("%s:%s", app.config.App.GrpcListenerHost, app.config.App.GrpcPort),
	)

	return grpcServer, listen, err
}
