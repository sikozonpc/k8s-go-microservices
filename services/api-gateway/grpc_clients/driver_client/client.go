package driver_client

import (
	"context"
	"os"

	pb "github.com/sikozonpc/ride-sharing/shared/proto/driver"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type DriverServiceClient struct {
	Client pb.DriverServiceClient
	conn   *grpc.ClientConn
}

func NewDriverServiceClient() (*DriverServiceClient, error) {
	driverServiceURL := os.Getenv("DRIVER_SERVICE_URL")
	if driverServiceURL == "" {
		driverServiceURL = "driver-service:9092"
	}

	conn, err := grpc.NewClient(driverServiceURL, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	client := pb.NewDriverServiceClient(conn)
	return &DriverServiceClient{
		Client: client,
		conn:   conn,
	}, nil
}

func (ds *DriverServiceClient) Close() {
	if ds.conn != nil {
		err := ds.conn.Close()
		if err != nil {
			return
		}
	}
}

func (ds *DriverServiceClient) StreamNearbyDrivers(ctx context.Context) (pb.DriverService_FindNearbyDriversClient, error) {
	return ds.Client.FindNearbyDrivers(ctx)
}
