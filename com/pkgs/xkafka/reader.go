package xkafka

import (
	"context"
	"errors"
	"fmt"
	"lark/com/pkgs/xlog"
	"lark/com/utils"
	"lark/pb"
	"time"

	"github.com/segmentio/kafka-go"
)

type Reader struct {
	readerGroup *kafka.Reader //用于广播
	readerSrv   *kafka.Reader //用于特点server消息
	cancle      context.CancelFunc
}

var reader *Reader

func NewReader(serverName string, serverId uint32, cfg *pb.Jkafka) *Reader {
	if reader != nil {
		return reader
	}
	reader = &Reader{
		readerGroup: kafka.NewReader(kafka.ReaderConfig{
			Brokers:        []string{cfg.Address},
			Partition:      int(cfg.Partition),
			Topic:          serverName,
			CommitInterval: time.Duration(cfg.ReadTimeOut) * time.Second,
			StartOffset:    kafka.FirstOffset,
			GroupID:        fmt.Sprintf("%s_%d", serverName, serverId),
		}),
		readerSrv: kafka.NewReader(kafka.ReaderConfig{
			Brokers:        []string{cfg.Address},
			Topic:          fmt.Sprintf("%s_%d", serverName, serverId),
			CommitInterval: time.Duration(cfg.ReadTimeOut) * time.Second,
			StartOffset:    kafka.FirstOffset,
		}),
	}
	return reader
}

// 启动Reader
func RunReader(callback func(msgid uint32, data []byte)) error {
	if reader == nil {
		return errors.New(utils.Const_Kafka_Reader_Not_Instance)
	}
	ctx, cancle := context.WithCancel(context.Background())
	reader.cancle = cancle
	fn := func(r *kafka.Reader) {
		for {
			if message, err := r.ReadMessage(ctx); err != nil {
				xlog.Warn(err.Error())
				break
			} else {
				key := utils.StrToUint32(string(message.Key))
				callback(key, message.Value)
			}
		}
	}
	go fn(reader.readerGroup)
	go fn(reader.readerSrv)
	return nil
}

// 退出reader
func ExitReader() {
	if reader == nil {
		return
	}
	reader.cancle()
	reader.readerGroup.Close()
	reader.readerSrv.Close()
}
