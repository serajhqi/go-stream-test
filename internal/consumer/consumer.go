package consumer

import (
	"fmt"

	"gitea.com/logicamp/lc"
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"go.uber.org/zap"
)

type message struct {
	Msg string
}

func (msg *message) String() string {
	return msg.Msg
}

type Feeder struct {
	producer    *lc.KafkaProducer
	consumer    *lc.KafkaConsumer
	FeedChannel chan any
}

func NewFeeder(kafkaHost, groupId string) *Feeder {
	feeder := &Feeder{
		producer:    lc.NewKafkaProducer(kafkaHost),
		consumer:    lc.NewKafkaConsumer(kafkaHost, groupId),
		FeedChannel: make(chan any),
	}
	feeder.init()

	return feeder
}

func (f *Feeder) init() {
	go func() {

		err := f.consumer.Consume(f.onMessage, []string{"test-stream"}, "earliest", false)
		if err != nil {
			lc.Logger.Error("consume error", zap.Error(err))
		}
	}()
}

func (f *Feeder) onMessage(message *kafka.Message) error {

	value := string(message.Value)
	fmt.Println("value is:", value)
	f.FeedChannel <- value

	return nil
}
