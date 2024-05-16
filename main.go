package main

import (
	"fmt"
	"lucianaUserServer/controller"
)

func main() {
	fmt.Println("正在运行")
	errCh1 := make(chan error)
	errCh2 := make(chan error)

	go controller.Register(errCh1)
	go controller.User(errCh2)

	for {
		select {
		case err := <-errCh1:
			fmt.Printf("注册服务挂掉了, %s\n", err)
		case err := <-errCh2:
			fmt.Printf("用户服务挂掉了, %s\n", err)
		}
	}
}
