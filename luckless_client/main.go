package main

import (
	"context"
	"io"
	"log"
	"os"
	"time"

	pb "github.com/luckless-finance/luckless/luckless"
	"google.golang.org/grpc"
)

const (
	// "server" is service name from docker-compose
	address     = "server:50052"
	defaultName = "world"
)

func main() {
	log.Printf("starting client...")
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	client := pb.NewMarketDataClient(conn)

	// Contact the server and print out its response.
	name := defaultName
	if len(os.Args) > 1 {
		name = os.Args[1]
	}

	// demo 1 call to Query
	{
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()

		r, err := client.Query(ctx, &pb.RangeRequest{
			Symbol:   name,
			First:    nil,
			Last:     nil,
			Calendar: pb.Calendar_COMPLETE,
		})
		if err != nil {
			log.Fatalf("could not Query: %v", err)
		}
		log.Printf("Query response: %s", r.GetData())
	}

	// demo 1 call to QueryStream
	{
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()
		stream, err := client.QueryStream(ctx, &pb.RangeRequest{
			Symbol:   name,
			First:    nil,
			Last:     nil,
			Calendar: pb.Calendar_COMPLETE,
		})

		if err != nil {
			log.Fatalf("%v.QueryStream(_) = _, %v", client, err)
		}

		for {
			dataPoint, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatalf("%v.QueryStream(_) = _, %v", client, err)
			}
			log.Printf("DataPoint: timestamp: %v double: %v", dataPoint.GetTimestamp(), dataPoint.Double)
		}
	}
}
