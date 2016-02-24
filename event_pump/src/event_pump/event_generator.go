package main

import (
	"math/rand"
	"time"

	uuid "github.com/satori/go.uuid"
)

var types = []string{
	"request",
	"delivered",
	"deferred",
	"click",
	"open",
	"processed",
	"dropped",
	"bounce",
	"spam_report",
	"unsubscribe",
	// "group_unsubscribe",
	// "group_resubscribe",
}

type Event struct {
	MessageID string `json:"sg_message_id"`
	Timestamp int64  `json:"timestamp"`
	Type      string `json:"type"`
}

func NewEventGenerator() *EventGenerator {
	g := &EventGenerator{
		ch: make(chan Event),
	}

	go g.start()

	return g
}

type EventGenerator struct {
	ch chan Event
}

func (g *EventGenerator) Generate() <-chan Event {
	return g.ch
}

func (g *EventGenerator) start() {
	for {
		e := Event{
			MessageID: uuid.NewV4().String(),
			Timestamp: time.Now().Unix(),
			Type:      types[0],
		}

		g.ch <- e

		go func() {
			done := false
			i := 0

			for !done {
				r := rand.Intn(10)
				<-time.After(time.Duration(r) * time.Second)

				e.Timestamp = time.Now().Unix()
				e.Type = randomType()

				g.ch <- e

				i++
				done = float64(1.0/(rand.Intn(len(types)-i)+1)) > .6
			}
		}()
	}
}

func randomType() string {
	// not request type
	return types[rand.Intn(len(types)-1)+1]
}
