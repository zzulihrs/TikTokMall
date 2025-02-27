package notify

import (
	"github.com/tiktokmall/backend/rpc_gen/kitex_gen/email"

	"github.com/kr/pretty"
)

type NoopEmail struct{}

func (e *NoopEmail) Send(req *email.EmailReq) error {
	// 处理，只是打印一下 req
	pretty.Printf("%v\n", req)
	return nil
}

func NewNoopEmail() NoopEmail {
	return NoopEmail{}
}
