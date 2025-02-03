package email

import (
	"log"
	"testing"

	"github.com/tiktokmall/backend/app/email/infra/mq"

	"github.com/nats-io/nats.go"
)

func TestConsumerInit(t *testing.T) {
	mq.Init()
	msg := &nats.Msg{Subject: "email", Data: []byte("Hello World, msg"), Header: make(nats.Header)}
	log.Printf("send msg: %v", msg)
	mq.Nc.PublishMsg(msg)
	mq.Nc.Publish("email", []byte("Hello World, msg"))
}
