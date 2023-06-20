package kafka

import (
	"context"
	"encoding/json"
	"strings"
	"time"

	env "github.com/amirnajdi/order-book/Helper"
	orderModel "github.com/amirnajdi/order-book/Models"
	"github.com/segmentio/kafka-go"
)

type OrderRequest struct {
	Order_id int
	Side     string
	Symbol   string
	Amount   float64
	Price    float64
}

func getBrokersAddress() []string {
	return strings.Split(env.Getenv("KAFKA_BROKERS_ADDRESS", "localhost:9092"), ",")
}

func Connection() *kafka.Reader {
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:         getBrokersAddress(),
		GroupID:         env.Getenv("KAFKA_GROUP_ID", "order_group"),
		Topic:           env.Getenv("KAFKA_TOPIC", "orders"),
		MinBytes:        10e3,            // 10KB
		MaxBytes:        10e6,            // 10MB
		MaxWait:         1 * time.Second, // Maximum amount of time to wait for new data to come when fetching batches of messages from kafka.
		ReadLagInterval: -1,
	})

	return reader
}

func ConsumeOrder(reader *kafka.Reader) (orderModel.Order, error) {
	message, err := reader.ReadMessage(context.Background())
	if err != nil {
		return orderModel.Order{}, err
	}

	return convertMessageToOrderType(message.Value)
}

func convertMessageToOrderType(messageValue []byte) (orderModel.Order, error) {
	var request OrderRequest
	if er := json.Unmarshal(messageValue, &request); er != nil {
		return orderModel.Order{}, er
	}

	order := orderModel.Order{
		ID: uint(request.Order_id),
		Side: request.Side,
		Symbol: request.Symbol,
		Amount: request.Amount,
		Price: request.Price,
	}

	return order, nil
}
