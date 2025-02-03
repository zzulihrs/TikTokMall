package rpc

import (
	"sync"

	"github.com/tiktokmall/backend/common/clientsuite"
	"github.com/tiktokmall/backend/rpc_gen/kitex_gen/cart/cartservice"
	"github.com/tiktokmall/backend/rpc_gen/kitex_gen/order/orderservice"
	"github.com/tiktokmall/backend/rpc_gen/kitex_gen/payment/paymentservice"
	"github.com/tiktokmall/backend/rpc_gen/kitex_gen/product/productcatalogservice"

	"github.com/tiktokmall/backend/app/checkout/conf"
	checkoutUtils "github.com/tiktokmall/backend/app/checkout/utils"

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
