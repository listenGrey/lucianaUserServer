package kafka

import (
	"encoding/json"
	"github.com/segmentio/kafka-go"
	"lucianaUserServer/dao"
	"lucianaUserServer/model"
	"time"

	"context"
)

// Register 用户注册
func Register() error {
	ctx := context.Background()
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:        []string{"localhost:9092"},
		Topic:          "register",
		CommitInterval: 1 * time.Second,
		GroupID:        "rec_team",
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

		err = dao.Register(&user)
		if err != nil {
			return err
		}
	}
}
