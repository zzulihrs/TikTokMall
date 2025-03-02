package mq

import (
	"fmt"
	"log"
	"os"
	"testing"
	"time"

	"github.com/nats-io/nats.go"
	"github.com/tiktokmall/backend/app/merchant/conf"
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

	go func() {

		Nc.Subscribe("testCount", func(msg *nats.Msg) {
			data := msg.Data
			fmt.Printf("%v\n", string(data))
		})

	}()
	time.Sleep(5 * time.Second)
}

func TestNats(t *testing.T) {
	os.Chdir("../../")
	// 空的话，是 127.0.0.1:4222
	Nc, err = nats.Connect("")
	log.Println("empty string", Nc.ConnectedAddr())
	Nc, err = nats.Connect(conf.GetConf().Nats.Address)
	log.Println("config string", Nc.ConnectedAddr())
}
