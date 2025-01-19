package email

import (
	email "bytetech/course/rpc_gen/kitex_gen/email"
	"context"
	"github.com/cloudwego/kitex/client/callopt"
	"github.com/cloudwego/kitex/pkg/klog"
)

func Send(ctx context.Context, req *email.EmailReq, callOptions ...callopt.Option) (resp *email.EmailResp, err error) {
	resp, err = defaultClient.Send(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "Send call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}
