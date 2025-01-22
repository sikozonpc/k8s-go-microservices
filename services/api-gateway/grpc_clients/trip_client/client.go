package trip_client

import (
	"context"
	"os"

	pb "github.com/sikozonpc/ride-sharing/shared/proto/trip"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type TripServiceClient struct {
	Client pb.TripServiceClient
	conn   *grpc.ClientConn
}

func NewTripServiceClient() (*TripServiceClient, error) {
	tripServiceURL := os.Getenv("TRIP_SERVICE_URL")
	if tripServiceURL == "" {
		tripServiceURL = "trip-service:9093"
	}

	conn, err := grpc.NewClient(tripServiceURL, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	client := pb.NewTripServiceClient(conn)
	return &TripServiceClient{
		Client: client,
		conn:   conn,
	}, nil
}

func (ds *TripServiceClient) Close() {
	if ds.conn != nil {
		err := ds.conn.Close()
		if err != nil {
			return
		}
	}
}

func (ds *TripServiceClient) CreateTrip(ctx context.Context, req *pb.CreateTripRequest) (*pb.CreateTripResponse, error) {
	return ds.Client.CreateTrip(ctx, req)
}
