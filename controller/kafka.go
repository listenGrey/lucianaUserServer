package controller

import (
	"encoding/json"
	"fmt"
	"github.com/segmentio/kafka-go"
	"lucianaUserServer/dao"
	"lucianaUserServer/model"
	"time"

	"context"
)

// RegisterService 用户注册
func RegisterService(address string) error {
	ctx := context.Background()
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:        []string{address},
		Topic:          "register",
		CommitInterval: 1 * time.Second,
		GroupID:        "rec_team",
		StartOffset:    kafka.FirstOffset,
	})

	fmt.Println(" kafka 服务正在运行")

	for {
		ms, err := reader.ReadMessage(ctx)
		if err != nil {
			return err
		}
		var user model.User
		err = json.Unmarshal(ms.Value, &user)
		if err != nil {
			return err
		}

		err = dao.Register(&user)
		if err != nil {
			return err
		}
	}
}
