package controllers

import (
	"context"
	"time"

	pb "github.com/DataDog/watermarkpodautoscaler/externalscaler"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type registeredClients struct {
	conn   *grpc.ClientConn
	client pb.ExternalScalerClient
}

// ExternalServerManager manages the connections to external gRPC metrics servers
type ExternalServerManager struct {
	connections map[string]*registeredClients
	fake        bool
}

// NewExternalServerManager creates a new ExternalServerManager
func NewExternalServerManager() *ExternalServerManager {
	return &ExternalServerManager{
		connections: make(map[string]*registeredClients),
		fake:        false,
	}
}

// SetFake sets the fake flag a nd allows the manager to be used for testing
func (e *ExternalServerManager) SetFake() *ExternalServerManager {
	e.fake = true
	return e
}

func (e *ExternalServerManager) registerConn(serverAddress string, conn *grpc.ClientConn, client pb.ExternalScalerClient) {
	e.connections[serverAddress] = &registeredClients{
		conn:   conn,
		client: client,
	}
}

// GetExternalMetric fetches the metric from the external gRPC server
func (e *ExternalServerManager) GetExternalMetric(metricName string, serverAddress string, metadata map[string]string) ([]int64, time.Time, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	if e.fake {
		return []int64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, time.Now(), nil
	}

	if _, ok := e.connections[serverAddress]; !ok {
		conn, err := grpc.Dial(serverAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			return nil, time.Now(), err
		}

		client := pb.NewExternalScalerClient(conn)
		e.registerConn(serverAddress, conn, client)
	}

	resp, err := e.connections[serverAddress].client.GetMetrics(ctx, &pb.GetMetricsRequest{
		MetricName: metricName,
		Metadata:   metadata,
	})

	if err != nil {
		return nil, time.Now(), err
	}

	metrics := make([]int64, 0)
	metrics = append(metrics, resp.Values...)

	return metrics, time.Now(), nil
}
