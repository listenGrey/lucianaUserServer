package main

import (
	"log"
	"lucianaUserServer/controller"
)

func main() {
	if err := controller.UserService("localhost:8964"); err != nil {
		log.Fatalf("Failed to connect, %s", err)
	}

	if err := controller.RegisterService("localhost:9092"); err != nil {
		log.Fatalf("kafka is not available, %s", err)
	}
}
