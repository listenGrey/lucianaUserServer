package controller

import (
	"fmt"
	"lucianaUserServer/logic"
)

func Register(errCh chan<- error) {
	fmt.Println("注册服务正在运行")
	for {
		if err := logic.RegisterService(); err != nil {
			errCh <- err
			return
		}
	}
}

func User(errCh chan<- error) {
	fmt.Println("用户服务正在运行")
	if err := logic.UserService(); err != nil {
		errCh <- err
		return
	}
}
