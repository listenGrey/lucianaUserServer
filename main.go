package main

import (
	"github.com/listenGrey/lucianagRpcPKG/user"
	"google.golang.org/grpc"
	"log"
	"net"

	service "lucianaUserServer/grpc"
	register "lucianaUserServer/kafka"
)

func main() {
	listen, err := net.Listen("tcp", "localhost:8964") //local ip and port
	if err != nil {
		log.Fatalf("Failed to connect, %s", err)
	}

	//初始化 gRpc server
	server := grpc.NewServer()

	user.RegisterCheckExistServer(server, &service.CheckExistenceServer{})
	user.RegisterLoginCheckServer(server, &service.LoginServer{})

	if err := server.Serve(listen); err != nil {
		log.Fatalf("Failed to connect, %s", err)
	}

	if err = register.Register(); err != nil {
		log.Fatalf("kafka is not available, %s", err)
	}
}
