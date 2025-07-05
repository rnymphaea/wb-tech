package kafka

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"github.com/segmentio/kafka-go"
	"wb-tech-l0/internal/database"
	"wb-tech-l0/internal/database/models"
)

type Consumer struct {
	reader  *kafka.Reader
	storage *database.Storage
}

func NewConsumer(brokers []string, topic string, groupID string, storage *database.Storage) *Consumer {
	return &Consumer{
		reader: kafka.NewReader(kafka.ReaderConfig{
			Brokers:  brokers,
			GroupID:  groupID,
			Topic:    topic,
			MinBytes: 10e3, // 10KB
			MaxBytes: 10e6, // 10MB
		}),
		storage: storage,
	}
}

func (c *Consumer) Run(ctx context.Context) {
	log.Println("Starting Kafka consumer...")
	
	for {
		select {
		case <-ctx.Done():
			log.Println("Stopping Kafka consumer")
			c.reader.Close()
			return
		default:
			msg, err := c.reader.FetchMessage(ctx)
			if err != nil {
				if ctx.Err() != nil {
					return
				}
				log.Printf("Error fetching message: %v", err)
				continue
			}

			start := time.Now()
			if err := c.processMessage(ctx, msg); err != nil {
				log.Printf("Error processing message: %v", err)
			} else {
				log.Printf("Processed message in %v (offset %d)", time.Since(start), msg.Offset)
			}
		}
	}
}

func (c *Consumer) processMessage(ctx context.Context, msg kafka.Message) error {
	var order models.Order
	if err := json.Unmarshal(msg.Value, &order); err != nil {
		log.Printf("Failed to unmarshal message: %v", err)
		return err
	}

	log.Printf("Received order: %s", order.UID)

	if err := c.storage.SaveOrder(ctx, &order); err != nil {
		log.Printf("Failed to save order %s: %v", order.UID, err)
		return err
	}

	log.Printf("Successfully saved order: %s", order.UID)
	return nil
}
