package rpc

import (
	"bytetech/course/common/clientsuite"
	"bytetech/course/rpc_gen/kitex_gen/cart/cartservice"
	"bytetech/course/rpc_gen/kitex_gen/order/orderservice"
	"bytetech/course/rpc_gen/kitex_gen/payment/paymentservice"
	"bytetech/course/rpc_gen/kitex_gen/product/productcatalogservice"
	"sync"

	"bytetech/course/app/checkout/conf"
	checkoutUtils "bytetech/course/app/checkout/utils"

	"github.com/cloudwego/kitex/client"
)

// client 进行服务发现

var (
	CartClient    cartservice.Client
	ProductClient productcatalogservice.Client
	PaymentClient paymentservice.Client
	OrderClient   orderservice.Client
	once          sync.Once
	opts          []client.Option
	err           error
	registryAddr  string
	serviceName   string
)

func Init() {
	once.Do(func() {
		registryAddr = conf.GetConf().Registry.RegistryAddress[0]
		serviceName = conf.GetConf().Kitex.Service
		opts = []client.Option{
			client.WithSuite(clientsuite.CommonServerSuite{
				RegistryAddr:       registryAddr,
				CurrentServiceName: serviceName,
			}),
		}
		initCartClient()
		initPaymentClient()
		initProductClient()
		initOrderClient()
	})
}

func initCartClient() {
	CartClient, err = cartservice.NewClient("cart", opts...)
	checkoutUtils.MustHandleError(err)
}

func initPaymentClient() {
	PaymentClient, err = paymentservice.NewClient("payment", opts...)
	checkoutUtils.MustHandleError(err)
}

func initProductClient() {
	ProductClient, err = productcatalogservice.NewClient("product", opts...)
	checkoutUtils.MustHandleError(err)
}

func initOrderClient() {
	OrderClient, err = orderservice.NewClient("order", opts...)
	checkoutUtils.MustHandleError(err)
}
