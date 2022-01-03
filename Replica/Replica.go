package main

import (
	"flag"
	"fmt"
	increment "github.com/Hexfall/DISYSMockExam/Increment"
	"google.golang.org/grpc"
	"log"
	"net"
)

var port = flag.Int("port", 5080, "port")
var target = flag.String("target", "", "target")

var server Server

func main() {
	flag.Parse()
	server = Server{
		isLeader:       *target == "",
		leaderAddr:     *target,
		selfAddr:       fmt.Sprintf(":%d", *port),
		value:          -1,
		replicaConns:   make(map[string]*grpc.ClientConn),
		replicaClients: make(map[string]increment.IncrementServiceClient),
	}

	// gRPC set-up.
	log.Printf("Attempting to listen on port: %d\n", *port)
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("Failed to listen on port %d. Error: %v", *port, err)
	} else {
		log.Printf("Now listening on port %d.\n", *port)
	}

	var options []grpc.ServerOption
	grpcServer := grpc.NewServer(options...)
	increment.RegisterIncrementServiceServer(grpcServer, &server)

	if *target != "" {
		go server.JoinCluster()
	}

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve gRPC client on port %d. Error: %v", *port, err)
	}
}
