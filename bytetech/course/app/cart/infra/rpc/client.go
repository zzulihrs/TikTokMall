package rpc

import (
	"bytetech/course/common/clientsuite"
	"bytetech/course/rpc_gen/kitex_gen/product/productcatalogservice"
	"sync"

	"bytetech/course/app/cart/conf"
	cartUtils "bytetech/course/app/cart/utils"

	"github.com/cloudwego/kitex/client"
)

// client 进行服务发现

var (
	ProductClient productcatalogservice.Client
	once          sync.Once
	err           error
	registryAddr  string
	serviceName   string
)

func Init() {
	once.Do(func() {
		registryAddr = conf.GetConf().Registry.RegistryAddress[0]
		serviceName = conf.GetConf().Kitex.Service
		initProductClient()
	})
}

func initProductClient() {
	opts := []client.Option{
		client.WithSuite(clientsuite.CommonServerSuite{
			RegistryAddr:       registryAddr,
			CurrentServiceName: serviceName,
		}),
	}
	ProductClient, err = productcatalogservice.NewClient("product", opts...)
	cartUtils.MustHandleError(err)
}
