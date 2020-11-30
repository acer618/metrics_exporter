package main

import (
    "context"
	"log"
	v1server "github.com/open-telemetry/opentelemetry-proto/gen/go/collector/metrics/v1"
)

type metricsServiceServer struct {
}

func NewMetricsServiceServer() v1server.MetricsServiceServer {
	return &metricsServiceServer{}
}

func (s *metricsServiceServer) Export(ctx context.Context, req *v1server.ExportMetricsServiceRequest) (*v1server.ExportMetricsServiceResponse, error) {
		log.Printf("Export invoked from client")
		response := &v1server.ExportMetricsServiceResponse{}
	    return response, nil
}
