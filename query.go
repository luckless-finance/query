package query

import (
	"context"
	"fmt"
	"github.com/golang/protobuf/ptypes"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
)

const (
	defaultPort = "50052"
)

type server struct {
	UnimplementedMarketDataServer
}

//https://github.com/grpc/grpc-go/blob/master/examples/route_guide/server/server.go
func (s *server) Query(ctx context.Context, in *RangeRequest) (*TimeSeries, error) {
	log.Printf("Received: %s", in.GetSymbol())
	dataPoint := DataPoint{
		Timestamp: ptypes.TimestampNow(),
		Double:    10,
	}
	fmt.Println(dataPoint.GetTimestamp())
	series := &TimeSeries{Data: []*DataPoint{
		&dataPoint, &dataPoint,
	}}
	fmt.Println(series)
	return series, nil
}

func (s *server) QueryStream(in *RangeRequest, stream MarketData_QueryStreamServer) error {
	log.Printf("Received: %s", in.GetSymbol())
	for _, v := range []int{1, 2, 3, 4} {
		dataPoint := DataPoint{
			Timestamp: ptypes.TimestampNow(),
			Double:    float64(v),
		}
		if err := stream.Send(&dataPoint); err != nil {
			return err
		}
	}
	return nil
}

func StartQueryListener() {
	log.Print("query server is starting up")
	var PORT = os.Getenv("PORT")
	if PORT == "" {
		PORT = defaultPort
		log.Printf("PORT defaulting to %s\n", PORT)
	}

	log.Printf("starting to listen on %s", PORT)
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", PORT))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	queryServer := grpc.NewServer()
	RegisterMarketDataServer(queryServer, &server{})
	if err := queryServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
