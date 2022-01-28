package main

import (
	"github.com/sulton0011/home_work/microservice/config"
	pb "github.com/sulton0011/home_work/microservice/genproto/users"
	"github.com/sulton0011/home_work/microservice/pkg/db"
	"github.com/sulton0011/home_work/microservice/pkg/logger"
	"github.com/sulton0011/home_work/microservice/service"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
    cfg := config.Load()

    log := logger.New(cfg.LogLevel, "template-service")
    defer logger.Cleanup(log)

    log.Info("main: sqlxConfig",
        logger.String("host", cfg.PostgresHost),
        logger.Int("port", cfg.PostgresPort),
        logger.String("database", cfg.PostgresDatabase))

    connDB, err := db.ConnectToDB(cfg)
    if err != nil {
        log.Fatal("sqlx connection to postgres error", logger.Error(err))
    }

    userService := service.NewUserService(connDB, log)

    lis, err := net.Listen("tcp", cfg.RPCPort)
    if err != nil {
        log.Fatal("Error while listening: %v", logger.Error(err))
    }

    s := grpc.NewServer()
    pb.RegisterUserServiceServer(s, userService)
    reflection.Register(s)
    log.Info("main: server running",
        logger.String("port", cfg.RPCPort))

    if err := s.Serve(lis); err != nil {
        log.Fatal("Error while listening: %v", logger.Error(err))
    }
}
