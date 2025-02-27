package email

import (
	"fmt"
	"log"
	"os"
	"testing"
	"time"

	"github.com/tiktokmall/backend/app/email/infra/mq"
	"github.com/tiktokmall/backend/rpc_gen/kitex_gen/email"
	"google.golang.org/protobuf/proto"
)

func TestConsumerInit(t *testing.T) {
	os.Chdir("../../../")
	mq.Init()
	go ConsumerInit()
	time.Sleep(time.Second)
	go func() {
		count := 1
		for {
			req := email.EmailReq{
				Content: fmt.Sprintf("%v", count),
			}
			data, err := proto.Marshal(&req)
			if err != nil {
				log.Printf("marshal err: %v", err)
				continue
			}
			log.Printf("publish msg, %+v", req)
			mq.Nc.Publish("email", data)
			count++
			time.Sleep(time.Second)
		}
	}()
	time.Sleep(10 * time.Second)
}
