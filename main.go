package main

import (
	"fmt"
	"lucianaUserServer/controller"
)

func main() {
	fmt.Println("用户服务正在运行")
	err := controller.UserService()
	if err != nil {
		fmt.Printf("用户服务挂掉了, %s\n", err)
	}
}
