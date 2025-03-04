package mq

import (
	"fmt"
	"log"
	"testing"
	"time"

	"github.com/nats-io/nats.go"
)

func TestPublishAndScribe(t *testing.T) {
	Nc, err = nats.Connect("nats://127.0.0.1:4222")
	if err != nil {
		log.Fatal(err)
	}
	go func() {
		for i := 0; i < 10; i++ {
			Nc.Publish("testCount", []byte(fmt.Sprintf("%v", i)))
			time.Sleep(time.Millisecond)
		}
	}()

	Nc.Subscribe("testCount", func(msg *nats.Msg) {
		data := msg.Data
		fmt.Printf("%v\n", string(data))
	})
	time.Sleep(5 * time.Second)
}
