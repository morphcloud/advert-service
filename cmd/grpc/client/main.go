package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/morphcloud/advert-service/internal/pb"
	"google.golang.org/grpc"
)

func main() {
	l := log.New(os.Stdout, "advert-client ", log.LstdFlags)

	conn, err := grpc.Dial("localhost:8081", grpc.WithInsecure())
	if err != nil {
		l.Fatalln(err)
	}

	defer conn.Close()

	client := pb.NewAdvertServiceClient(conn)

	stream, err := client.ReadAll(context.Background(), &empty.Empty{})
	if err != nil {
		l.Fatalln(err)
	}
	for {
		advert, err := stream.Recv()
		if err == io.EOF {
			break
		}

		if err != nil {
			l.Fatalln(err)
		}

		fmt.Println(advert)
	}
}
