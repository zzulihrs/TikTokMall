package rpc

import (
	"bytetech/course/common/clientsuite"
	"bytetech/course/rpc_gen/kitex_gen/cart/cartservice"
	"bytetech/course/rpc_gen/kitex_gen/checkout/checkoutservice"
	"bytetech/course/rpc_gen/kitex_gen/order/orderservice"
	"bytetech/course/rpc_gen/kitex_gen/product/productcatalogservice"
	"bytetech/course/rpc_gen/kitex_gen/user/userservice"
	"sync"

	"bytetech/course/app/frontend/conf"
	frontendUtils "bytetech/course/app/frontend/utils"

	"github.com/cloudwego/kitex/client"
)

// client 进行服务发现

var (
	UserClient     userservice.Client
	ProductClient  productcatalogservice.Client
	CartClient     cartservice.Client
	CheckoutClient checkoutservice.Client
	OrderClient    orderservice.Client
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
