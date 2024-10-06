package grpc

import (
	"context"
	"github.com/listenGrey/lucianagRpcPKG/user"
	"google.golang.org/grpc/peer"
	"log"
	"lucianaUserServer/dao"
	"lucianaUserServer/model"
	"lucianaUserServer/mq"
)

type CheckExistenceServer struct {
	user.UnimplementedCheckExistServer
}

func (u *CheckExistenceServer) CheckExist(ctx context.Context, email *user.Email) (*user.Exist, error) {
	_, ok := peer.FromContext(ctx)
	if ok {
		log.Printf("检查该用户是否存在")
	}

	// 检查邮箱
	exist, err := dao.CheckEmail(email.GetEmail())
	if err != nil {
		return nil, err
	}

	return &user.Exist{Exist: exist}, nil
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
	if err != nil {
		return nil, err
	}

	return model.LogInfoMarshal(info), nil
}

type RegisterServer struct {
	user.UnimplementedRegisterCheckServer
}

func (u *RegisterServer) RegisterCheck(ctx context.Context, re *user.RegisterFrom) (*user.Success, error) {
	_, ok := peer.FromContext(ctx)
	if ok {
		log.Printf("用户注册")
	}

	// 用户注册
	err := mq.RegisterQueue(re)
	if err != nil {
		return &user.Success{Success: false}, err
	}
	return &user.Success{Success: true}, nil
}
