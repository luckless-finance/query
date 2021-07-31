package main

import (
	"context"
	"fmt"
	"github.com/golang/protobuf/ptypes"
	query "github.com/luckless-finance/luckless"
	"google.golang.org/grpc"
	"log"
	"os"
	"time"
)

const (
	// "server" is service name from docker-compose
	defaultQueryHost = "localhost"
	queryPort        = "50052"
	defaultName      = "FOO"
	defaultSeries    = "CLOSE"
	defaultFirst     = "2020-01-01T00:00:00Z" // time.Date(2020, time.January, 1, 0, 0, 0, 0, time.UTC)
	defaultLast      = "2021-01-01T00:00:00Z" // time.Date(2020, time.January, 1, 0, 0, 0, 0, time.UTC)
)

func main() {
	queryHost := os.Getenv("QUERY_HOST")
	if queryHost == "" {
		queryHost = defaultQueryHost
		log.Printf("QUERY_HOST defaulting to %s\n", defaultQueryHost)
	}
	queryAddress := fmt.Sprintf("%s:%s", queryHost, queryPort)

	log.Printf("starting client to query server @ %s", queryAddress)
	// Set up a connection to the server.
	conn, err := grpc.Dial(queryAddress, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	client := query.NewMarketDataClient(conn)

	// Contact the server and print out its response.
	name := defaultName
	series := defaultSeries
	if len(os.Args) > 2 {
		name = os.Args[1]
		series = os.Args[2]
	}
	firstStr := defaultFirst
	lastStr := defaultLast
	if len(os.Args) > 4 {
		firstStr = os.Args[3]
		lastStr = os.Args[4]
	}

	// demo 1 call to Query
	{
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()
		firstTime, err := time.Parse(time.RFC3339, firstStr)
		if err != nil {
			log.Fatalf("Could not parse first date: %s\n", firstStr)
		}
		first, err := ptypes.TimestampProto(firstTime)
		if err != nil {
			log.Fatalf("Could not build first TimestampProto")
		}
		lastTime, err := time.Parse(time.RFC3339, lastStr)
		if err != nil {
			log.Fatalf("Could not parse last date: %s\n", lastStr)
		}
		last, err := ptypes.TimestampProto(lastTime)
		if err != nil {
			log.Fatalf("Could not build TimestampProto")
		}
		r, err := client.Query(ctx, &query.RangedRequest{
			Symbol: name,
			Series: series,
			First:  first,
			Last:   last,
		})
		if err != nil {
			log.Fatalf("could not Query: %v", err)
		}
		log.Printf("Query response: %s", r.GetData())
	}
}
