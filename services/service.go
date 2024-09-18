package services

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/segmentio/kafka-go"
	"log"
	"os"
	"strconv"
	"taxi_order_service/models"
)

type LocationService struct {
	writer *kafka.Writer
}

func (p *LocationService) StoreLocation(ctx context.Context, point models.Point, userId int) error {
	var pointJson, err = json.Marshal(point)
	userIdString := strconv.Itoa(userId)
	if err != nil {
		return fmt.Errorf("could not marshal point: %w", err)
	}
	err = p.WriteMessage(ctx, []byte(userIdString), pointJson)
	if err != nil {
		return fmt.Errorf("could not write message to kafka: %w", err)
	}
	return nil
}

func NewLocationService(brokers []string, topic string) *LocationService {
	return &LocationService{
		writer: &kafka.Writer{
			Addr:     kafka.TCP(brokers...),
			Topic:    topic,
			Balancer: &kafka.Hash{},
			Logger:   log.New(os.Stdout, "kafka writer: ", 0),
		},
	}
}

func (p *LocationService) WriteMessage(ctx context.Context, key, value []byte) error {
	msg := kafka.Message{
		Key:   key,
		Value: value,
	}

	err := p.writer.WriteMessages(ctx, msg)
	if err != nil {
		return fmt.Errorf("could not write message to Kafka: %w", err)
	}
	return nil
}

func (p *LocationService) Close() error {
	return p.writer.Close()
}
