package main

import (
	"fmt"
	"log"
	"lucianaUserServer/controller"
)

func main() {
	fmt.Println("正在运行")
	if err := controller.UserService(":8964"); err != nil {
		log.Fatalf("Failed to connect, %s", err)
	}

	if err := controller.RegisterService(":9092"); err != nil {
		log.Fatalf("kafka is not available, %s", err)
	}
}
