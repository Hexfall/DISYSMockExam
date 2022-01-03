package main

import (
	"context"
	increment "github.com/Hexfall/DISYSMockExam/Increment"
	"google.golang.org/grpc"
	"log"
)

type FrontEnd struct {
	client  increment.IncrementServiceClient
	conn    *grpc.ClientConn
	ctx     context.Context
	backups []string
}

func (fe *FrontEnd) Increment() int64 {
	mes, err := fe.client.Increment(fe.ctx, &increment.VoidMessage{})
	if err != nil {
		log.Fatalf("Failed to increment value. Error: %v", err)
	}
	return mes.Number
}

func (fe *FrontEnd) Connect(ip string) {
	fe.conn.Close()

	var options []grpc.DialOption
	options = append(options, grpc.WithBlock(), grpc.WithInsecure())
	conn, err := grpc.Dial(ip, options...)
	if err != nil {
		log.Fatalf("Failed to dial gRPC server on ip %s. Error: %v", ip, err)
	}
	fe.conn = conn

	fe.ctx = context.Background()
	fe.client = increment.NewIncrementServiceClient(fe.conn)

	// Find the leader of the cluster.
	mes, err := fe.client.GetLeader(fe.ctx, &increment.VoidMessage{})
	if err != nil {
		log.Fatalf("Failed to get leader. Error %v", err)
	}
	if !mes.IsHost {
		fe.Connect(mes.Ip)
	}
}
