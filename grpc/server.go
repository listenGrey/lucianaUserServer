package grpc

import (
	"context"
	"github.com/listenGrey/lucianagRpcPKG/user"
	"google.golang.org/grpc/peer"
	"log"
	"lucianaUserServer/dao"
	"lucianaUserServer/model"
)

type ExistenceServer struct {
	user.UnimplementedCheckExistenceServer
}

func (u *ExistenceServer) RegisterCheck(ctx context.Context, email *user.RegisterEmail) (*user.Existence, error) {
	_, ok := peer.FromContext(ctx)
	if ok {
		log.Printf("检查该用户是否存在")
	}

	// 检查邮箱
	exist, err := dao.CheckEmail(email.GetEmail())
	if err != nil {
		return &user.Existence{Exist: exist, ServerError: true}, nil
	}

	return &user.Existence{Exist: exist, ServerError: false}, nil
}

type RegisterServer struct {
	user.UnimplementedRegisterInfoServer
}

func (r *RegisterServer) Register(ctx context.Context, form *user.RegisterForm) (*user.Success, error) {
	_, ok := peer.FromContext(ctx)
	if ok {
		log.Printf("用户注册")
	}

	// 用户注册
	info, err := dao.Register(form)
	if err != nil {
		return &user.Success{Success: info, ServerError: true}, nil
	}

	return &user.Success{Success: info, ServerError: false}, nil
}

type LoginServer struct {
	user.UnimplementedLoginCheckServer
}

func (u *LoginServer) LoginCheck(ctx context.Context, user *user.LoginForm) (*user.LogInfo, error) {
	_, ok := peer.FromContext(ctx)
	if ok {
		log.Printf("用户登录")
	}

	// 用户登录
	info, err := dao.Login(user)
	serverErr := false
	if err != nil {
		serverErr = true
	}

	return model.LogInfoMarshal(info, serverErr), nil
}
