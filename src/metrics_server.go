package main

import (
	"log"
	"net"
	"google.golang.org/grpc"

	"github.com/open-telemetry/opentelemetry-proto/gen/go/collector/metrics/v1"
)

func main() {
	lis, err := net.Listen("tcp", ":9999")
	if err != nil {
		log.Fatalf("failed to listen on port 9999: %v", err)
	}

	s := metricsServiceServer{}

	grpcServer := grpc.NewServer()
	v1.RegisterMetricsServiceServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve gRPC server over port 9999: %v", err);
	}

}
