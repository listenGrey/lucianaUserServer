package logic

import (
	"encoding/json"
	"github.com/segmentio/kafka-go"
	"lucianaUserServer/conf"
	"lucianaUserServer/dao"
	"lucianaUserServer/model"
	"time"

	"context"
)

// RegisterService 用户注册
func RegisterService() error {
	ctx := context.Background()

	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:        []string{conf.KafkaServerAddress},
		Topic:          "register",
		CommitInterval: 1 * time.Second,
		GroupID:        "register",
		StartOffset:    kafka.FirstOffset,
	})

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

		// 处理消息
		err = dao.Register(&user)
		if err != nil {
			return err
		}
	}
}
