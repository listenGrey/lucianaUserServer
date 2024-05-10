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

	// 创建 10 个 Kafka 消费者组，每个消费者组一个协程
	for i := 0; i < 10; i++ {
		go func(groupID string) {
			reader := kafka.NewReader(kafka.ReaderConfig{
				Brokers:        []string{address},
				Topic:          "register",
				CommitInterval: 1 * time.Second,
				GroupID:        groupID,
				StartOffset:    kafka.FirstOffset,
			})

			fmt.Printf("Consumer with GroupID %s is running\n", groupID)

			for {
				ms, err := reader.ReadMessage(ctx)
				if err != nil {
					fmt.Printf("Error reading message for GroupID %s: %s\n", groupID, err)
					continue
				}

				var user model.User
				err = json.Unmarshal(ms.Value, &user)
				if err != nil {
					fmt.Printf("Error unmarshalling message for GroupID %s: %s\n", groupID, err)
					continue
				}

				// 处理消息
				err = dao.Register(&user)
				if err != nil {
					fmt.Printf("Error registering user for GroupID %s: %s\n", groupID, err)
					continue
				}

				// 手动提交消息的偏移量
				if err = reader.CommitMessages(ctx, ms); err != nil {
					fmt.Printf("Error committing offset for GroupID %s: %s\n", groupID, err)
				}
			}
		}(fmt.Sprintf("register_%d", i+1))
	}

	// 阻塞主 goroutine
	select {}
}
