package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"strings"

	env "github.com/amirnajdi/order-book/Helper"
	"github.com/schollz/progressbar/v3"
	"github.com/segmentio/kafka-go"
)

type OrderRequest struct {
	Order_id int
	Side     string
	Symbol   string
	Amount   float64
	Price    int
}

var side = []string{"BUY", "SELL"}
var symbol = []string{"BTC", "USDT", "HEX", "TRX", "SHIB"}

const minimumNumberOfData int = 1000

func main() {
	address := strings.Split(env.Getenv("KAFKA_BROKERS_ADDRESS", "localhost:9092"), ",")
	writer := &kafka.Writer{
		Addr:     kafka.TCP(address...),
		Topic:    env.Getenv("KAFKA_TOPIC", "orders"),
		Balancer: &kafka.LeastBytes{},
	}

	var numberOfTestData int
	fmt.Printf("How much test data do you want to insert on Kafka? (minimum: %d): ", minimumNumberOfData)
	fmt.Scanln(&numberOfTestData)
	if numberOfTestData < minimumNumberOfData {
		log.Fatal("The number must be bigger than the minimum data")
	}

	bar := progressbar.Default(int64(numberOfTestData))
	
	var messages []kafka.Message
	var counter int
	for i := 1; i <= numberOfTestData; i++ {
		counter++
		
		// make fake order 
		order := OrderRequest{
			Order_id: rand.Intn(10000000),
			Price:    rand.Intn(100000000),
			Amount:   (rand.Float64() * float64(rand.Intn(100))) + 1,
			Symbol:   symbol[rand.Intn(len(symbol))],
			Side:     side[rand.Intn(len(side))],
		}

		// convert fake order to json for push on Kafka
		orderJson, convertToJsonError := json.Marshal(order)
		if convertToJsonError != nil {
			fmt.Println(convertToJsonError)
			continue
		}

		messages = append(messages, kafka.Message{
			Key:   []byte(string(rune(i))),
			Value: orderJson,
		})

		// to increase performance push every 100 order together in Kafka
		if counter == minimumNumberOfData {
			if err := writer.WriteMessages(context.Background(), messages...); err != nil {
				log.Fatal("failed to write messages:", err)
			}

			messages = []kafka.Message{}
			counter = 0
			bar.Add(minimumNumberOfData)			
		}
	}

	if err := writer.Close(); err != nil {
		log.Fatal("failed to close writer:", err)
	}
}
