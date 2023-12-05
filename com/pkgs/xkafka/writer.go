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
	"google.golang.org/protobuf/proto"
)

type Writer struct {
	writer     *kafka.Writer
	retryCount int //重试次数
}

var writer *Writer

func NewWriter(cfg *pb.Jkafka) *Writer {
	if writer != nil {
		return writer
	}
	writer.writer = &kafka.Writer{
		Addr:                   kafka.TCP(cfg.Address),
		WriteTimeout:           time.Duration(cfg.WriteTimeOut) * time.Second,
		RequiredAcks:           kafka.RequireOne,
		AllowAutoTopicCreation: cfg.AllowAutoTopic,
	}
	writer.retryCount = int(cfg.WriteRetry)
	return writer
}

// 发送给特定server（server id为0，则广播）
func SendToServer(serverName string, serverId uint32, msgid uint32, msg proto.Message) error {
	if writer == nil {
		return errors.New(utils.Const_Kafka_Writer_Not_Instance)
	}
	dat, err := proto.Marshal(msg)
	if err != nil {
		return err
	}
	topic := serverName
	if serverId > 0 {
		topic = fmt.Sprintf("%s_%d", serverName, serverId)
	}
	return writer.send(topic, msgid, dat)
}

// 广播给所有该类型的server
func BroadcastToServer(serverName string, msgid uint32, msg proto.Message) error {
	return SendToServer(serverName, 0, msgid, msg)
}

func (t *Writer) send(topic string, msgid uint32, data []byte) (err error) {
	key := []byte(utils.ToString(msgid))
	for i := 0; i < t.retryCount; i++ {
		err = t.writer.WriteMessages(context.Background(),
			kafka.Message{
				Topic: topic,
				Key:   key,
				Value: data,
			},
		)
		if err != nil {
			if err != kafka.LeaderNotAvailable {
				xlog.Warn(err.Error())
			}
		} else {
			break
		}
	}
	return
}

func ExitWriter() {
	if writer != nil {
		writer.writer.Close()
	}
}
