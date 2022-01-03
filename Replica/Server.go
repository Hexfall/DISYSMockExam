package main

import (
	"context"
	increment "github.com/Hexfall/DISYSMockExam/Increment"
	"log"
	"sync"
)

type Server struct {
	increment.UnimplementedIncrementServiceServer

	isLeader     bool
	leaderAddr   string
	selfAddr     string
	replicas     []string
	replicaMutex sync.Mutex

	value int64
	mutex sync.Mutex
}

// gRPC functions.

func (s *Server) Increment(ctx context.Context, void *increment.VoidMessage) (*increment.IncrementMessage, error) {
	if !s.isLeader {
		// This replica doesn't have the authority to Increment.
		return nil, &increment.ImpermissibleError{}
	}

	s.mutex.Lock()
	defer s.mutex.Unlock()

	s.value++
	log.Printf("Value has been incremented to %d.\n", s.value)
	return &increment.IncrementMessage{Number: s.value}, nil
}

func (s *Server) GetLeader(ctx context.Context, void *increment.VoidMessage) (*increment.LeaderMessage, error) {
	return &increment.LeaderMessage{
		Ip:       s.leaderAddr,
		IsLeader: s.isLeader,
	}, nil
}

func (s *Server) GetReplicas(ctx context.Context, void *increment.VoidMessage) (*increment.ReplicaListMessage, error) {
	return &increment.ReplicaListMessage{Ips: s.replicas}, nil
}

func (s *Server) Join(ctx context.Context, ipMessage *increment.IpMessage) (*increment.VoidMessage, error) {
	if !s.isLeader {
		return nil, &increment.ImpermissibleError{}
	}

	s.replicaMutex.Lock()
	defer s.replicaMutex.Unlock()

	s.replicas = append(s.replicas, ipMessage.Ip)
	// TODO: Monitor heartbeat.

	return &increment.VoidMessage{}, nil
}

func (s *Server) HeartBeat(ctx context.Context, void *increment.VoidMessage) (*increment.VoidMessage, error) {
	return &increment.VoidMessage{}, nil
}

func (s *Server) SendReplicas(ctx context.Context, replicasMessage *increment.ReplicaListMessage) (*increment.VoidMessage, error) {
	if s.isLeader {
		// Leader cannot be ordered around.
		return nil, &increment.ImpermissibleError{}
	}
	// Should really check whether sender is leader, somehow.
	s.replicas = replicasMessage.Ips

	return &increment.VoidMessage{}, nil
}

func (s *Server) SendValue(ctx context.Context, incrementMessage *increment.IncrementMessage) (*increment.VoidMessage, error) {
	if s.isLeader {
		// Leader cannot be ordered around.
		return nil, &increment.ImpermissibleError{}
	}
	// Should really check whether sender is leader, somehow.
	s.value = incrementMessage.Number

	return &increment.VoidMessage{}, nil
}
