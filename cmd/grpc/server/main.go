package main

import (
	"log"
	"net"
	"os"
	"os/signal"

	"github.com/joho/godotenv"
	"github.com/morphcloud/advert-service/internal/impl"
	"github.com/morphcloud/advert-service/internal/pb"
	"google.golang.org/grpc"
)

func main() {
	l := log.New(os.Stdout, "advert-service ", log.LstdFlags)

	if err := godotenv.Load(); err != nil {
		l.Fatalln(err)
	}

	port := os.Getenv("PORT")
	netListener := getNetListener(port)
	grpcServer := grpc.NewServer()

	advertServiceServer := impl.NewAdvertServiceServer()

	pb.RegisterAdvertServiceServer(grpcServer, advertServiceServer)

	l.Println("Server is running on port " + port)

	go func() {
		if err := grpcServer.Serve(netListener); err != nil {
			l.Fatalln(err)
		}
	}()

	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)

	sig := <-sigChan
	l.Println("Received terminate. Graceful shutdown...", sig)

	grpcServer.GracefulStop()
}

func getNetListener(port string) net.Listener {
	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalln(err)
	}
	return lis
}
