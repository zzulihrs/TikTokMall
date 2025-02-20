package rpc

import (
	"sync"

	"github.com/tiktokmall/backend/common/clientsuite"
	"github.com/tiktokmall/backend/rpc_gen/kitex_gen/cart/cartservice"
	"github.com/tiktokmall/backend/rpc_gen/kitex_gen/checkout/checkoutservice"
	"github.com/tiktokmall/backend/rpc_gen/kitex_gen/merchant/merchantservice"
	"github.com/tiktokmall/backend/rpc_gen/kitex_gen/order/orderservice"
	"github.com/tiktokmall/backend/rpc_gen/kitex_gen/payment/paymentservice"
	"github.com/tiktokmall/backend/rpc_gen/kitex_gen/product/productcatalogservice"
	"github.com/tiktokmall/backend/rpc_gen/kitex_gen/user/userservice"

	"github.com/tiktokmall/backend/app/frontend/conf"
	frontendUtils "github.com/tiktokmall/backend/app/frontend/utils"

	"github.com/cloudwego/kitex/client"
)

// client 进行服务发现

var (
	UserClient     userservice.Client
	ProductClient  productcatalogservice.Client
	CartClient     cartservice.Client
	CheckoutClient checkoutservice.Client
	OrderClient    orderservice.Client
	PaymentClient  paymentservice.Client
	MerchantClient merchantservice.Client
	once           sync.Once
	opts           []client.Option
	err            error
	registryAddr   string
)

func Init() {
	once.Do(func() {
		registryAddr = conf.GetConf().Hertz.RegistryAddr
		opts = []client.Option{
			client.WithSuite(clientsuite.CommonServerSuite{
				RegistryAddr:       registryAddr,
				CurrentServiceName: frontendUtils.ServiceName,
			}),
		}
		initUserClient()
		initProductClient()
		initCartClient()
		initCheckoutClient()
		initOrderClient()
		initPaymentClient()
		initMerchantClient()
	})
}

func initUserClient() {
	UserClient, err = userservice.NewClient("user", opts...)
	frontendUtils.MustHandleError(err)
}

func initProductClient() {
	ProductClient, err = productcatalogservice.NewClient("product", opts...)
	frontendUtils.MustHandleError(err)
}

func initCartClient() {
	CartClient, err = cartservice.NewClient("cart", opts...)
	frontendUtils.MustHandleError(err)
}

func initCheckoutClient() {
	CheckoutClient, err = checkoutservice.NewClient("checkout", opts...)
	frontendUtils.MustHandleError(err)
}

func initOrderClient() {
	OrderClient, err = orderservice.NewClient("order", opts...)
	frontendUtils.MustHandleError(err)
}

func initPaymentClient() {
	PaymentClient, err = paymentservice.NewClient("payment", opts...)
	frontendUtils.MustHandleError(err)
}
func initMerchantClient() {
	MerchantClient, err = merchantservice.NewClient("merchant", opts...)
	frontendUtils.MustHandleError(err)
}
