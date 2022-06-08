package main

import (
	"context"
	"log"
	"net"
	"sync"
	pb "thinknetica_golang_core/Lesson_14-RPC/messages_proto"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type Signalman struct {
	mu   sync.Mutex
	Data []pb.Message

	// Композиция интерфейса.
	pb.SignalmanServer
}

func (s *Signalman) Messages(_ *pb.Empty, stream pb.Signalman_MessagesServer) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	for _, item := range s.Data {
		stream.Send(&item)
	}
	return nil
}

func (s *Signalman) Send(_ context.Context, message *pb.Message) (*pb.Empty, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.Data = append(s.Data, *message)
	return nil, nil
}

func main() {
	log.Println("Server starting")
	srv := Signalman{}
	messages1 := pb.Message{Id: 1, Text: "The Go Programming Language", CreatedAt: timestamppb.Now()}
	messages2 := pb.Message{Id: 2, Text: "1984", CreatedAt: timestamppb.Now()}
	srv.Data = append(srv.Data, messages1, messages2)

	lis, err := net.Listen("tcp", ":12345")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterSignalmanServer(grpcServer, &srv)
	grpcServer.Serve(lis)
}
