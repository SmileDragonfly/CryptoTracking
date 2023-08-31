package main

import (
	grpc_telesvc "github.com/SmileDragonfly/go-lib/crypto-grpc/telegram-service"
	"google.golang.org/grpc"
	"net"
	"telesvc/logger"
)

func main() {
	// Init log
	err := logger.NewLogger("./config/logcfg.json")
	if err != nil {
		panic(err)
	}
	// Init server
	lis, err := net.Listen("tcp", "127.0.0.1:3001")
	if err != nil {
		logger.Instance.Error(err)
		return
	}
	newServer := grpc.NewServer()
	grpc_telesvc.RegisterTelegramServiceServer(newServer, &Server{})
	err = newServer.Serve(lis)
	if err != nil {
		logger.Instance.Error(err)
		return
	}
}
