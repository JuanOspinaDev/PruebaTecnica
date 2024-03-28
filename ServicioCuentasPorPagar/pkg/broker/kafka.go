package broker

import (
	"context"
	"log"

	"github.com/segmentio/kafka-go"
)

// KafkaConfig almacena la configuración para conectar con Kafka
type KafkaConfig struct {
	Brokers []string
	Topic   string
}

type KafkaPublisher struct {
	writer *kafka.Writer
}

func NewKafkaPublisher(cfg KafkaConfig) *KafkaPublisher {
	return &KafkaPublisher{
		writer: &kafka.Writer{
			Addr:     kafka.TCP(cfg.Brokers...),
			Topic:    cfg.Topic,
			Balancer: &kafka.LeastBytes{},
		},
	}
}

func (p *KafkaPublisher) Publish(ctx context.Context, key, value []byte) error {
	return p.writer.WriteMessages(ctx,
		kafka.Message{
			Key:   key,
			Value: value,
		},
	)
}

func (p *KafkaPublisher) Close() error {
	return p.writer.Close()
}


type KafkaConsumer struct {
	reader *kafka.Reader
}

func NewKafkaConsumer(cfg KafkaConfig, groupID string) *KafkaConsumer {
	return &KafkaConsumer{
		reader: kafka.NewReader(kafka.ReaderConfig{
			Brokers:  cfg.Brokers,
			Topic:    cfg.Topic,
			GroupID:  groupID,
			MinBytes: 10e3, // 10KB
			MaxBytes: 10e6, // 10MB
		}),
	}
}

func (c *KafkaConsumer) StartListening(ctx context.Context, handleMessage func(kafka.Message) error) {
	for {
		msg, err := c.reader.ReadMessage(ctx)
		if err != nil {
			log.Printf("error reading message from kafka: %v", err)
			continue
		}
		if err := handleMessage(msg); err != nil {
			log.Printf("error handling message: %v", err)
		}
	}
}

func (c *KafkaConsumer) Close() error {
	return c.reader.Close()
}
