package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"

	"github.com/golang/protobuf/ptypes"

	pb "github.com/luckless-finance/luckless/luckless"

	"google.golang.org/grpc"
)

const (
	defaultPort = "50052"
)

type server struct {
	pb.UnimplementedMarketDataServer
}

//https://github.com/grpc/grpc-go/blob/master/examples/route_guide/server/server.go
func (s *server) Query(ctx context.Context, in *pb.RangeRequest) (*pb.TimeSeries, error) {
	log.Printf("Received: %s", in.GetSymbol())
	dataPoint := pb.DataPoint{
		Timestamp: ptypes.TimestampNow(),
		Double:    10,
	}
	fmt.Println(dataPoint.GetTimestamp())
	series := &pb.TimeSeries{Data: []*pb.DataPoint{
		&dataPoint, &dataPoint,
	}}
	fmt.Println(series)
	return series, nil
}

func (s *server) QueryStream(in *pb.RangeRequest, stream pb.MarketData_QueryStreamServer) error {
	log.Printf("Received: %s", in.GetSymbol())
	for _, v := range []int{1, 2, 3, 4} {
		dataPoint := pb.DataPoint{
			Timestamp: ptypes.TimestampNow(),
			Double:    float64(v),
		}
		if err := stream.Send(&dataPoint); err != nil {
			return err
		}
	}
	return nil
}

func main() {
	log.Print("query server is starting up")
	var PORT = os.Getenv("PORT")
	if PORT == "" {
		PORT = defaultPort
		log.Printf("PORT defaulting to %s\n", PORT)
	}

	log.Printf("starting to listen on %s", PORT)
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s",PORT))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	queryServer := grpc.NewServer()
	pb.RegisterMarketDataServer(queryServer, &server{})
	if err := queryServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
