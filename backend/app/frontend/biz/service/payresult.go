package service

import (
	"context"
	"fmt"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/tiktokmall/backend/app/frontend/infra/rpc"
	"github.com/tiktokmall/backend/rpc_gen/kitex_gen/order"
	"log"
	http "net/http"
)

const (
	// AppId
	appId = "9021000144642150"
	// 应用私钥，这是我自己申请的
	privateKey = "MIIEpAIBAAKCAQEAlLUQR8gCipxZCMKGxx4p9xeFunvjQmVbGkaRWzvicaYwzQ3PnrcVoeEJmm6WtuRoNgHHdsglkwFFJ4asBwGnnoJfpoWX4d1q25jbVT6lQDFqrirMAapnn6tACwFVF+xfZTEkXP6yWg3FHS3FrsCuWu12h9deQbaA3ooy5Of/oRLiJqhtj79BO4U7PcaCvqW5geCtvJjjZp1YteEcWvKI2YfwEyQNU16BDAp0DdJGmsakVl+ILtmKq4EWRnHpA4WKFEnSDlakeG5+ETdVjGuADWsAXIG/N6Y667IDKGMlRjTAX7eDcH/LJksJTpQmVqxndmhTH9DHCri70druTRFw+wIDAQABAoIBAQCCRyovKLuJQbYCpaE7+mIbdgETQgv0FCW/HEml2w3kMQuZ6VqWGqGaAzPNlsr1CZ9+iQp9NpUsd7VlDmfRb5KwnFk94tYP61H9dqnYwZw2Y9SpJxc48lf8GyRtP8qTveJJHcdUrVIE8QFaMcp48ZcOU/E+yuBpw4Cc9gfmYJWkPzfLcOJpVzE5my3Giz5gx6qCeY0/Vw3pSBgjqJ0QbrHPZZBHBF0nQJBoZMzoZicW3uFzDNgCyrHP7Gt5vpHS65Q4gxvYI1uevbDLYCp/zKl+zwPB9ULVu3gzqXI6t7WfTFUwiInz7A5/TBEsWl/oMTTrSSwHdVmonD2oFmbU+GkxAoGBAMoGwNTu9N+iee6Zamo/R0MEzrlkJRgtEu/NG/dFCqk76izsV4yvxTg2X8eKZJMv/qdMLLXsPyTvQIj1o/YOl4Ld3M3/umH2i3WXL4L+tCkU2KqsrhVtGvofsPWngEY96u7m2hUW1LT9CdKCB4WiY7WHvfioCzPGEILF7z58E+apAoGBALxvpGc8wopLO4BbJGyEtdy9pEZI489uriGtWENXsQMDb1kXwInzXy3tGd1/Gm5lNBOOkawgP+iVf/+bP3HDCM+0fK6vGk0eaRDPUEO8KrfHEE0c3IKm9QrE+on6Y0lekxdnbGmeEfdbwbRkQsWH9pz3CL1Vbb3U7OVCrI0QJvUDAoGBAJvW3cTjl20zLg+JtHbE3TmrabPEtoCRtHvaomn5jiHBoACLR3W02NNlzjhkXvTTHwL6Vbr3xDW/gO6lXZce5m0Cq6MUUzfiMO3Cc9n+lAbkl8YAckEA6sBq1dyJGwAUHzeuSCfgdrAuSPdjal4BSBzt2vMG3a1QacgW21g96jQRAoGAcYPvziFUMGtR1OkS8CyiQfAcXhra4cMTM0ZxvV++Sspu8YTVgEUUOV50DnLbQVXGIWHpb3+eAEbPbPPhLG8Jh9Z6peDmPz7qnC3HdIaOVVeeAlY4oJxjUbGIHEayOSi4A4lSTe3jdNfZwQoFD2nwrm5C3YvxKEJ000dvGyt3zHECgYBhOl09MmrGHn1XhkfdAOdzVHZsxNt6cHrf2fCfB2Bl4jsnZd+LJV1puXMpkYbG7WXiU6wnoPA76wNY9ovzKUInxLIsezBdpWCwdBAQl3MM6BejO3lig1+Jpiilg2uv3yWleavS/pq1id8cVzeZSPDK08S3GZ7c1IEOTjZdMUUX/A=="
)

type PayresultService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewPayresultService(Context context.Context, RequestContext *app.RequestContext) *PayresultService {
	return &PayresultService{RequestContext: RequestContext, Context: Context}
}

func (h *PayresultService) Run(req *http.Request) (resp map[string]any, err error) {
	// 需要修改订单状态---等待提供修改订单状态的接口
	// 解析异步通知的参数

	out_trade_no := h.RequestContext.Query("out_trade_no")
	log.Println("out_trade_no: ", out_trade_no)

	_, err = rpc.OrderClient.ChangeOrderStatus(h.Context, &order.ChangeOrderStatusReq{
		OrderId:     out_trade_no,
		OrderStatus: 2, // 订单状态 0 - 创建订单待支付 1 - 支付成功 2 - 支付失败 3 - 取消订单
	})

	if err != nil {
		return nil, fmt.Errorf("change order status failed: %v", err)
	}

	return map[string]any{
		"message": "Success",
	}, nil
}
