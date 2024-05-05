package controller

import (
	"github.com/listenGrey/lucianagRpcPKG/user"
	"google.golang.org/grpc"
	service "lucianaUserServer/grpc"
	"net"
)

func UserService(address string) error {
	listen, err := net.Listen("tcp", address) //local ip and port
	if err != nil {
		return err
	}

	//初始化 gRpc server
	server := grpc.NewServer()

	user.RegisterCheckExistServer(server, &service.CheckExistenceServer{})
	user.RegisterLoginCheckServer(server, &service.LoginServer{})

	if err = server.Serve(listen); err != nil {
		return err
	}

	return nil
}
