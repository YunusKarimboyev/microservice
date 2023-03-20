package main

import (
	"net"

	"github.com/double/test_microservice/config"
	u "github.com/double/test_microservice/genproto/user"
	"github.com/double/test_microservice/pkg/db"
	"github.com/double/test_microservice/pkg/logger"
	"github.com/double/test_microservice/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	cfg := config.Load()
	log := logger.New(cfg.LogLevel, "golang")
	defer logger.Cleanup(log)

	connDb, err := db.ConnectToDb(cfg)
	if err != nil {
		log.Fatal("Error connect postgres", logger.Error(err))
	}

	userService := service.NewUserService(connDb, log)

	lis, err := net.Listen("tcp", cfg.UserServicePort)

	if err != nil {
		log.Fatal("Error while listening: %v", logger.Error(err))
	}

	s := grpc.NewServer()
	reflection.Register(s)
	u.RegisterUserServiceServer(s, userService)

	log.Info("main: server running",
		logger.String("port", cfg.UserServicePort))
	if err := s.Serve(lis); err != nil {
		log.Fatal("Error while listening: %v", logger.Error(err))
	}

}
