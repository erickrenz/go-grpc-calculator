package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/erickrenz/go-grpc-calculator/pb"
)

func main() {
	serverAddr := flag.String(
		"server", "localhost:8080",
		"The server address is in the format of host:port",
	)
	flag.Parse()

	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	conn, err := grpc.DialContext(ctx, *serverAddr, opts...)
	if err != nil {
		log.Fatalln("fail to dial:", err)
	}
	defer conn.Close()

	client := pb.NewCalculatorClient(conn)

	res, err := client.Sum(ctx, &pb.NumbersRequest{
		Numbers: []int64{10, 9, 8, 7, 6},
	})
	if err != nil {
		log.Fatalln("error sending request:", err)
	}

	fmt.Println("result:", res.Result)
}
