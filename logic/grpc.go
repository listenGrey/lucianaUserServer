package logic

import (
	"fmt"
	"github.com/listenGrey/lucianagRpcPKG/user"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"lucianaUserServer/conf"
	service "lucianaUserServer/grpc"
	"net"
)

func UserService() error {
	creds, err := credentials.NewServerTLSFromFile(conf.CertFile, conf.KeyFile) // crt,key
	if err != nil {
		return err
	}
	listen, err := net.Listen("tcp", conf.GrpcServerAddress) //local ip and port
	if err != nil {
		return err
	}

	//初始化 gRpc server
	server := grpc.NewServer(grpc.Creds(creds))

	user.RegisterCheckExistServer(server, &service.CheckExistenceServer{})
	user.RegisterLoginCheckServer(server, &service.LoginServer{})

	if err = server.Serve(listen); err != nil {
		return err
	}
	fmt.Println(" gRpc 服务正在运行")
	return nil
}
