package main

import (
	"context"
	"fmt"
	"time"

	"github.com/segmentio/kafka-go"
)

const (
	topic     = "my-topic"
	partition = 0
)

func main() {
	//建立连接
	// conn, err := kafka.DialLeader(context.Background(), "tcp", "localhost:9092", topic, partition)
	// if err != nil {
	// 	log.Fatal("fail to dial leader", err.Error())
	// }
	// conn.SetWriteDeadline(time.Now().Add(10 * time.Second))
	// _, err = conn.WriteMessages(
	// 	kafka.Message{Value: []byte("One!")},
	// 	kafka.Message{Value: []byte("Two!")},
	// 	kafka.Message{Value: []byte("Three!")},
	// )
	// if err != nil {
	// 	log.Fatal("fail to write message", err.Error())
	// }
	// if err = conn.Close(); err != nil {
	// 	log.Fatal("fail to close writer", err.Error())
	// }
	writer(context.Background())
}

func writer(ctx context.Context) {
	writer := &kafka.Writer{
		Addr: kafka.TCP("localhost:9092"),
		// Topic:                  topic,
		Balancer:               &kafka.Hash{},
		WriteTimeout:           3 * time.Second,
		RequiredAcks:           kafka.RequireOne,
		AllowAutoTopicCreation: true,
	}
	var err error
	for i := 0; i < 3; i++ {
		err = writer.WriteMessages(ctx,
			kafka.Message{
				Topic: topic + "_1",
				Key:   []byte("1"),
				Value: []byte("The One!")},
		)
		if err != nil {
			if err == kafka.LeaderNotAvailable {
				continue
			} else {
				fmt.Println(err.Error())
			}
		} else {
			break
		}
	}
	for i := 0; i < 3; i++ {
		err = writer.WriteMessages(ctx,
			kafka.Message{
				Topic: topic,
				Key:   []byte("1"),
				Value: []byte("The Two!"),
			},
		)
		if err != nil {
			if err == kafka.LeaderNotAvailable {
				continue
			} else {
				fmt.Println(err.Error())
			}
		} else {
			break
		}
	}
	writer.Close()
}
