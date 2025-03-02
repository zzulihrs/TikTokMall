package order

import (
	"context"
	"github.com/tiktokmall/backend/app/order/biz/dal/mysql"
	"github.com/tiktokmall/backend/app/order/biz/model"
	"log"
	"time"

	"github.com/cloudwego/kitex/server"
	"github.com/nats-io/nats.go"
	"github.com/tiktokmall/backend/app/order/infra/mq"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
)

func ConsumerInit() {
	tracer := otel.Tracer("shop-nats-consumer")
	sub, err := mq.Nc.Subscribe("ORDER.CREATE", func(msg *nats.Msg) {

		orderID := string(msg.Data)
		go func() {
			delay := time.Duration(30*60) * time.Second
			time.Sleep(delay)
			if err := model.ChangeOrderStatus(mysql.DB, context.Background(), orderID, 3); err != nil {
				log.Println("mq 更新订单状态失败:", err)
				return
			}

		}()

		ctx := context.Background()
		ctx = otel.GetTextMapPropagator().Extract(ctx, propagation.HeaderCarrier(msg.Header))
		_, span := tracer.Start(ctx, "shop-nats-consumer")
		defer span.End()

		// log.Printf("subscribe msg, %v", req)
	})
	if err != nil {
		panic(err)
	}

	server.RegisterShutdownHook(func() {
		sub.Unsubscribe()
		mq.Nc.Close()
	})
}
