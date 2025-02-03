package mtl

import (
	"github.com/kitex-contrib/obs-opentelemetry/provider"
)

func InitTracing(serviceName string) provider.OtelProvider {
	p := provider.NewOpenTelemetryProvider(
		provider.WithServiceName(serviceName),
		// 默认值
		// provider.WithExportEndpoint("localhost:4317"),
		provider.WithInsecure(),
		provider.WithEnableMetrics(false),
	)
	return p

}
