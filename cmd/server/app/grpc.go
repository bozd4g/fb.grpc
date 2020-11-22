package app

import (
	"errors"
	"fmt"
	"github.com/bozd4g/fb.grpc/cmd/server/grpc/usergrpcservice"
	"github.com/bozd4g/fb.grpc/cmd/server/internal/application/userservice"
	"github.com/bozd4g/fb.grpc/cmd/server/internal/infrastructure/repository/userrepository"
	"github.com/bozd4g/fb.grpc/pkg/postgresql"
	pb "github.com/bozd4g/fb.grpc/pkg/proto/userservice"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"net"
	"os"
)

func NewGrpcApplication() IApplication {
	return &Application{}
}

func (application *Application) BuildGrpc() IApplication {
	application.logger = *logrus.New()
	application.AddPostgreSql(postgresql.Opts{
		Host:     "localhost",
		User:     "postgres",
		Password: "123456",
		Database: "fbgrpc",
		Port:     5432,
	})

	application.grpcserver = grpc.NewServer()

	userRepository := userrepository.New(application.db)
	userService := userservice.New(userRepository)
	pb.RegisterUserServiceServer(application.grpcserver, &usergrpcservice.UserGrpcService{
		Logger:      application.logger,
		UserService: userService,
	})

	return application
}

func (application *Application) RunGrpc() error {
	port := os.Getenv("GRPC_PORT")
	if port == "" {
		port = ":50051"
	}

	lis, err := net.Listen("tcp", port)
	if err != nil {
		message := fmt.Sprintf("failed to listen: %v", err)
		application.logger.Error(message)
		return errors.New(message)
	}

	fmt.Println(fmt.Sprintf("gRPC service listening and serving on %s", port))
	err = application.grpcserver.Serve(lis)
	if err != nil {
		message := fmt.Sprintf("failed to listen: %v", err)
		application.logger.Error(message)
		return errors.New(message)
	}

	return nil
}
