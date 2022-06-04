package main

import (
	"context"
	"fmt"
	"io"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/timestamppb"

	pb "thinknetica_golang_core/Lesson_14-RPC/messages_proto"
)

func main() {
	conn, err := grpc.Dial("localhost:12345", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()

	client := pb.NewSignalmanClient(conn)

	err = printAllMessagesOnserver(client)
	if err != nil {
		fmt.Println(err)
		return
	}

	client.Send(context.Background(), &pb.Message{Id: 3, Text: "The Lord Of The Rings", CreatedAt: timestamppb.Now()})

	err = printAllMessagesOnserver(client)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func printAllMessagesOnserver(client pb.SignalmanClient) error {
	fmt.Println("\nЗапрашиваю сообщения на gRPC-сервере.")
	stream, err := client.Messages(context.Background(), &pb.Empty{})
	if err != nil {
		return err
	}

	for {
		book, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		fmt.Printf("Id: %v, Text: %v, CreatedAt: %v\n", book.Id, book.Text, book.CreatedAt.AsTime())
	}
	return nil
}
