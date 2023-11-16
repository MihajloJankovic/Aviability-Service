package main

import (
	"context"
	"log"
	"net"
	"os"
	"time"

	"github.com/MihajloJankovic/Aviability-Service/handlers"
	protos "github.com/MihajloJankovic/Aviability-Service/protos/main"
	"google.golang.org/grpc"
)

func main() {

	lis, err := net.Listen("tcp", ":9095")
	if err != nil {
		log.Fatal(err)
	}
	serverRegistar := grpc.NewServer()

	timeoutContext, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	logger := log.New(os.Stdout, "[aviability-main] ", log.LstdFlags)
	log := log.New(os.Stdout, "[aviability-repo-log] ", log.LstdFlags)

	AviabilityRepo, err := handlers.New(timeoutContext, log)
	if err != nil {
		logger.Fatal(err)
	}
	defer AviabilityRepo.Disconnect(timeoutContext)

	// NoSQL: Checking if the connection was established
	AviabilityRepo.Ping()

	//Initialize the handler and inject said logger
	service := handlers.NewServer(logger, AviabilityRepo)

	protos.RegisterAccommodationAviabilityServer(serverRegistar, service)
	err = serverRegistar.Serve(lis)
	if err != nil {
		log.Fatal(err)
	}
}
