package main

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/segmentio/kafka-go"
)

const (
	topic     = "my-topic"
	partition = 0
)

func main() {
	// conn, err := kafka.DialLeader(context.Background(), "tcp", "localhost:9092", topic, partition)
	// if err != nil {
	// 	log.Fatal("fail to dial leader", err.Error())
	// }
	// conn.SetReadDeadline(time.Now().Add(10 * time.Second))
	// // fetch 10kb min 1MB max
	// batch := conn.ReadBatch(10e3, 1e6)
	// b := make([]byte, 10e3)
	// for {
	// 	n, err := batch.Read(b)
	// 	if err != nil {
	// 		log.Fatal("fail to read message", err.Error())
	// 		break
	// 	}
	// 	log.Println(string(b[:n]))
	// }
	// if err = batch.Close(); err != nil {
	// 	log.Fatal("fail to close batch", err.Error())
	// }
	// if err = conn.Close(); err != nil {
	// 	log.Fatal("fail to close reader", err.Error())
	// }
	reader(context.Background())
}

func reader(ctx context.Context) {
	wg := sync.WaitGroup{}
	//用于广播
	readerGroup := kafka.NewReader(kafka.ReaderConfig{
		Brokers:        []string{"localhost:9092"},
		Topic:          topic,
		CommitInterval: 3 * time.Second,
		StartOffset:    kafka.FirstOffset,
		GroupID:        topic + "_1",
	})
	//用于发送特定server
	readerSrv := kafka.NewReader(kafka.ReaderConfig{
		Brokers:        []string{"localhost:9092"},
		Topic:          topic + "_1",
		CommitInterval: 3 * time.Second,
		StartOffset:    kafka.FirstOffset,
	})
	wg.Add(2)
	go func() {
		defer wg.Done()
		for {
			if message, err := readerGroup.ReadMessage(ctx); err != nil {
				fmt.Println(err.Error())
				break
			} else {
				fmt.Println(string(message.Value))
			}
		}
	}()
	go func() {
		defer wg.Done()
		for {
			if message, err := readerSrv.ReadMessage(ctx); err != nil {
				fmt.Println(err.Error())
				break
			} else {
				fmt.Println(string(message.Value))
			}
		}
	}()
	wg.Wait()
	if err := readerGroup.Close(); err != nil {
		log.Fatal("reader close error", err.Error())
	}
	if err := readerSrv.Close(); err != nil {
		log.Fatal("reader close error", err.Error())
	}
}
