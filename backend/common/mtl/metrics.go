package mtl

import (
	"net"
	"net/http"

	"github.com/cloudwego/kitex/pkg/registry"
	"github.com/cloudwego/kitex/server"
	consul "github.com/kitex-contrib/registry-consul"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/collectors"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var Registry *prometheus.Registry

func InitMetric(serviceName, metricsPort, registryAddr string) (registry.Registry, *registry.Info) {
	Registry = prometheus.NewRegistry()
	// go 运行时相关的指标
	Registry.MustRegister(collectors.NewGoCollector())
	// 进程相关的指标
	Registry.MustRegister(collectors.NewProcessCollector(collectors.ProcessCollectorOpts{}))
	r, _ := consul.NewConsulRegister(registryAddr)
	addr, _ := net.ResolveTCPAddr("tcp", metricsPort)
	registryInfo := &registry.Info{
		ServiceName: "prometheus",
		Addr:        addr,
		Weight:      10,
		Tags:        map[string]string{"service": serviceName},
	}
	_ = r.Register(registryInfo)
	server.RegisterShutdownHook(func() {
		r.Deregister(registryInfo)
	})
	http.Handle("/metrics", promhttp.HandlerFor(Registry, promhttp.HandlerOpts{}))
	go http.ListenAndServe(metricsPort, nil)
	return r, registryInfo
}
