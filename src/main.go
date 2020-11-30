package main

import (
	"context"
	"log"
	"google.golang.org/grpc"

	"github.com/open-telemetry/opentelemetry-proto/gen/go/collector/metrics/v1"
)

func main() {
	sendMetrics(context.Background())
}

func sendMetrics(ctx context.Context) {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":9999", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %s", err)
	}

	defer conn.Close()

	c := v1.NewMetricsServiceClient(conn)
	request := v1.ExportMetricsServiceRequest {
	}

	response, err := c.Export(ctx, &request)
	if err != nil {
		log.Fatalf("error sending metrics to server: %v", err)
	} else {
		log.Printf("Response from metrics server: %s", response)
	}
}



