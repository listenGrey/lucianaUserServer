package mq

import (
	"encoding/json"
	"fmt"
	"github.com/listenGrey/lucianagRpcPKG/user"
	"github.com/segmentio/kafka-go"
	"golang.org/x/net/context"
	"lucianaUserServer/conf"
	"lucianaUserServer/model"
)

// RegisterQueue 用户注册
func RegisterQueue(r *user.RegisterFrom) error {
	ctx := context.Background()
	// 创建 Kafka 生产者
	writer := &kafka.Writer{
		Addr:  kafka.TCP(conf.KafkaServerAddress),
		Topic: "register",
		//Balancer:               &kafka.Hash{},
		//WriteTimeout:           1 * time.Second,
		//RequiredAcks:           kafka.RequireNone,
		AllowAutoTopicCreation: false,
	}

	defer writer.Close()

	// 构造消息
	key := []byte(fmt.Sprintf("%d", r.Id))             // key = id
	value, err := json.Marshal(model.UserUnmarshal(r)) // value = data
	if err != nil {
		return err
	}

	// 发送消息
	err = writer.WriteMessages(
		ctx,
		kafka.Message{
			Key:   key,
			Value: value,
		},
	)
	if err != nil {
		return err
	}

	return nil
}