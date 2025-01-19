package main

import (
	"github.com/nats-io/nats.go"
)

func main() {
	nc, _ := nats.Connect(nats.DefaultURL)
	defer nc.Close()
	// Simple Async Subscriber
	nc.Publish("foo", []byte("Hello World"))
	msg := &nats.Msg{Subject: "foo", Data: []byte("Hello World, msg"), Header: make(nats.Header)}
	nc.PublishMsg(msg)
}
