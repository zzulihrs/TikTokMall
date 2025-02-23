package service

import (
	"context"
	"github.com/cloudwego/kitex/pkg/kerrors"
	"github.com/go-pay/gopay"
	myalipay "github.com/go-pay/gopay/alipay"
	payment "github.com/tiktokmall/backend/rpc_gen/kitex_gen/payment"
	"strconv"
)

const (
	// AppId
	appId = "9021000144642150"
	// 应用私钥，这是我自己申请的
	privateKey = "MIIEpAIBAAKCAQEAlLUQR8gCipxZCMKGxx4p9xeFunvjQmVbGkaRWzvicaYwzQ3PnrcVoeEJmm6WtuRoNgHHdsglkwFFJ4asBwGnnoJfpoWX4d1q25jbVT6lQDFqrirMAapnn6tACwFVF+xfZTEkXP6yWg3FHS3FrsCuWu12h9deQbaA3ooy5Of/oRLiJqhtj79BO4U7PcaCvqW5geCtvJjjZp1YteEcWvKI2YfwEyQNU16BDAp0DdJGmsakVl+ILtmKq4EWRnHpA4WKFEnSDlakeG5+ETdVjGuADWsAXIG/N6Y667IDKGMlRjTAX7eDcH/LJksJTpQmVqxndmhTH9DHCri70druTRFw+wIDAQABAoIBAQCCRyovKLuJQbYCpaE7+mIbdgETQgv0FCW/HEml2w3kMQuZ6VqWGqGaAzPNlsr1CZ9+iQp9NpUsd7VlDmfRb5KwnFk94tYP61H9dqnYwZw2Y9SpJxc48lf8GyRtP8qTveJJHcdUrVIE8QFaMcp48ZcOU/E+yuBpw4Cc9gfmYJWkPzfLcOJpVzE5my3Giz5gx6qCeY0/Vw3pSBgjqJ0QbrHPZZBHBF0nQJBoZMzoZicW3uFzDNgCyrHP7Gt5vpHS65Q4gxvYI1uevbDLYCp/zKl+zwPB9ULVu3gzqXI6t7WfTFUwiInz7A5/TBEsWl/oMTTrSSwHdVmonD2oFmbU+GkxAoGBAMoGwNTu9N+iee6Zamo/R0MEzrlkJRgtEu/NG/dFCqk76izsV4yvxTg2X8eKZJMv/qdMLLXsPyTvQIj1o/YOl4Ld3M3/umH2i3WXL4L+tCkU2KqsrhVtGvofsPWngEY96u7m2hUW1LT9CdKCB4WiY7WHvfioCzPGEILF7z58E+apAoGBALxvpGc8wopLO4BbJGyEtdy9pEZI489uriGtWENXsQMDb1kXwInzXy3tGd1/Gm5lNBOOkawgP+iVf/+bP3HDCM+0fK6vGk0eaRDPUEO8KrfHEE0c3IKm9QrE+on6Y0lekxdnbGmeEfdbwbRkQsWH9pz3CL1Vbb3U7OVCrI0QJvUDAoGBAJvW3cTjl20zLg+JtHbE3TmrabPEtoCRtHvaomn5jiHBoACLR3W02NNlzjhkXvTTHwL6Vbr3xDW/gO6lXZce5m0Cq6MUUzfiMO3Cc9n+lAbkl8YAckEA6sBq1dyJGwAUHzeuSCfgdrAuSPdjal4BSBzt2vMG3a1QacgW21g96jQRAoGAcYPvziFUMGtR1OkS8CyiQfAcXhra4cMTM0ZxvV++Sspu8YTVgEUUOV50DnLbQVXGIWHpb3+eAEbPbPPhLG8Jh9Z6peDmPz7qnC3HdIaOVVeeAlY4oJxjUbGIHEayOSi4A4lSTe3jdNfZwQoFD2nwrm5C3YvxKEJ000dvGyt3zHECgYBhOl09MmrGHn1XhkfdAOdzVHZsxNt6cHrf2fCfB2Bl4jsnZd+LJV1puXMpkYbG7WXiU6wnoPA76wNY9ovzKUInxLIsezBdpWCwdBAQl3MM6BejO3lig1+Jpiilg2uv3yWleavS/pq1id8cVzeZSPDK08S3GZ7c1IEOTjZdMUUX/A=="
	// 域名，在在这里我测试用的是本机暴露给公网的地址，后续部署了记得换成服务器的地址
	serverDomin = "https://381e-103-172-183-54.ngrok-free.app"
	NotifyUrl   = "/payresult"
	ReturnUrl   = "/paysuccess"
)

type AlipayService struct {
	ctx context.Context
} // NewAlipayService new AlipayService
func NewAlipayService(ctx context.Context) *AlipayService {
	return &AlipayService{ctx: ctx}
}

// Run create note info
func (s *AlipayService) Run(req *payment.AlipayReq) (resp *payment.AlipayResp, err error) {
	//初始化支付宝客户端
	//    appId：应用ID
	//    privateKey：应用私钥，支持PKCS1和PKCS8
	//    isProd：是否是正式环境,false是沙箱环境
	client, err := myalipay.NewClient(appId, privateKey, false)
	if err != nil {
		return nil, kerrors.NewBizStatusError(40005, "ailipay initialized falied")
	}
	//配置公共参数
	client.SetCharset("utf-8").
		SetSignType(myalipay.RSA2).
		SetNotifyUrl(serverDomin + NotifyUrl).
		SetReturnUrl(serverDomin + ReturnUrl)

	//请求参数
	bm := make(gopay.BodyMap)
	bm.Set("subject", "网站测试支付")
	bm.Set("out_trade_no", req.TransactionId)
	bm.Set("total_amount", strconv.FormatFloat(float64(req.TotalAmount), 'f', 2, 64))
	bm.Set("product_code", "FAST_INSTANT_TRADE_PAY")

	ctx := context.Background()

	//电脑网站支付请求
	payUrl, err := client.TradePagePay(ctx, bm)
	if err != nil {
		return nil, kerrors.NewBizStatusError(40005, "TradePagePay falied")
	}
	return &payment.AlipayResp{
		PayUrl: payUrl,
	}, nil

}
