package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/shopify/sarama"
)

const MaxEvents = 1000

var Tick = time.Second

func main() {
	rand.Seed(time.Now().Unix())

	brokerList := os.Getenv("KAFKA_BROKER_LIST")
	brokers := strings.Split(brokerList, ",")

	topic := os.Getenv("KAFKA_TOPIC")

	log.Printf("using brokers: %v with topic %s", brokers, topic)

	run(brokers, topic)
}

func run(brokers []string, topic string) {
	producer, err := sarama.NewSyncProducer(brokers, nil)
	if err != nil {
		log.Fatalln(err)
	}

	defer func() {
		if err := producer.Close(); err != nil {
			log.Fatalln(err)
		}
	}()

	sigCh := make(chan os.Signal)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)
	accumulator := make(map[string]int)

	eg := NewEventGenerator()

	func() {
		for {
			select {
			case e := <-eg.Generate():
				b, err := json.Marshal(&e)
				if err != nil {
					log.Println("unable to json marshal", err)
					continue
				}

				msg := &sarama.ProducerMessage{
					Topic: topic,
					Value: sarama.ByteEncoder(b),
				}

				log.Printf("sending %s for %s", e.Type, e.MessageID)
				_, _, err = producer.SendMessage(msg)
				if err != nil {
					log.Printf("FAILED to send message: %s", err)
					continue
				}

				accumulator[e.Type]++

			case <-sigCh:
				log.Println("exiting")
				return
			}
		}
	}()

	for e, count := range accumulator {
		fmt.Println("%s: %d", e, count)
	}
}
