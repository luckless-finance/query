package query

import (
	"context"
	"fmt"
	"github.com/golang/protobuf/ptypes"
	"google.golang.org/grpc"
	"net"
	"os"
	"time"

	"github.com/withmandala/go-log"
)

var logger = log.New(os.Stdout).WithColor()

const (
	defaultPort = "50052"
)

type server struct {
	UnimplementedMarketDataServer
}

func (rr *RangedRequest) validateRangedRequest() error {
	if rr.GetSymbol() == "" {
		return fmt.Errorf("symbol empty")
	}
	if rr.GetSeries() == "" {
		return fmt.Errorf("series empty")
	}
	if rr.GetFirst() == nil {
		return fmt.Errorf("first nil")
	}
	if rr.GetLast() == nil {
		return fmt.Errorf("last nil")
	}
	logger.Debug("OK RangedRequest is valid")
	return nil
}

//https://github.com/grpc/grpc-go/blob/master/examples/route_guide/server/server.go
func (s *server) Query(ctx context.Context, in *RangedRequest) (*TimeSeries, error) {
	err := in.validateRangedRequest()
	if err != nil {
		fmt.Println(err)
	}
	logger.Infof("Received Query for symbol|series: '%s'|'%s'", in.GetSymbol(), in.GetSeries())
	first, err := ptypes.Timestamp(in.GetFirst())
	if err != nil {
		return nil, err
	}
	logger.Infof("query first: %s\n", first.Format(time.RFC3339))

	last, err := ptypes.Timestamp(in.GetLast())
	if err != nil {
		return nil, err
	}
	logger.Infof("query last: %s\n", last.Format(time.RFC3339))

	trig := NewTrig(10, 10, 0, 100)
	timeSeries, err := trig.timeSeriesPtrs(first, last)
	if err != nil {
		return nil, err
	}
	series := &TimeSeries{Data: timeSeries}
	logger.Infof("TimeSeries len: %d", len(series.Data))
	return series, nil
}

func StartQueryListener() {
	logger.Info("query server is starting up")
	var PORT = os.Getenv("PORT")
	if PORT == "" {
		PORT = defaultPort
		logger.Infof("PORT defaulting to %s\n", PORT)
	}

	logger.Infof("starting to listen on %s", PORT)
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", PORT))
	if err != nil {
		logger.Fatalf("failed to listen: %v", err)
	}
	queryServer := grpc.NewServer()
	RegisterMarketDataServer(queryServer, &server{})
	if err := queryServer.Serve(lis); err != nil {
		logger.Fatalf("failed to serve: %v", err)
	}
}
