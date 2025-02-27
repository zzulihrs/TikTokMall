package user

import (
	"context"
	"encoding/json"

	"github.com/tiktokmall/backend/app/user/infra/mq"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/server"
	"github.com/nats-io/nats.go"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
)

var (
	handlers map[string]func(*mq.UserMsg)
)

func ConsumerInit() {
	tracer := otel.Tracer("shop-nats-consumer")
	sub, err := mq.Nc.Subscribe("user_email", func(msg *nats.Msg) {
		var userMsg mq.UserMsg
		err := json.Unmarshal(msg.Data, &userMsg)
		if err != nil {
			klog.Error(err)
			return
		}

		ctx := context.Background()
		ctx = otel.GetTextMapPropagator().Extract(ctx, propagation.HeaderCarrier(msg.Header))
		_, span := tracer.Start(ctx, "shop-nats-consumer")
		defer span.End()

		// 根据 msg 更新 redis 中的 user 数据。目前仅删除
		// log.Printf("subscribe msg, %v", req)
		// noopEmail := notify.NewNoopEmail()
		// _ = noopEmail.Send(&req)
	})
	if err != nil {
		panic(err)
	}

	server.RegisterShutdownHook(func() {
		sub.Unsubscribe()
		mq.Nc.Close()
	})
}

func handleUserMsg(usermsg *mq.UserMsg) {
	if usermsg.Method == "Delete" {

	}
}

func DeleteCacheUser(data map[string]any) {
	// 根据 redis 的 key 进行删除
	// cacheKey := data["key"]

}
