package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

	"github.com/jhomoreira/grpc/pb"
	"google.golang.org/grpc"
)

func main() {
	connection, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect to gRPC Server: %v", err)
	}

	defer connection.Close()

	client := pb.NewUserServiceClient(connection)
	//AddUser(client)
	//AddUserVerbose(client)
	AddUsers(client)
}

func AddUser(client pb.UserServiceClient) {
	req := &pb.User{
		Id:    "0",
		Name:  "Jho",
		Email: "jho@flanp.com",
	}

	res, err := client.AddUser(context.Background(), req)
	if err != nil {
		log.Fatalf("Could not make gRPC request: %v", err)
	}

	fmt.Println(res)
}

func AddUserVerbose(client pb.UserServiceClient) {
	req := &pb.User{
		Id:    "0",
		Name:  "Jho",
		Email: "jho@flanp.com",
	}

	responseStream, err := client.AddUserVerbose(context.Background(), req)
	if err != nil {
		log.Fatalf("Could not make gRPC request: %v", err)
	}

	for {
		stream, err := responseStream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Coud not receive the msg: %v", err)
		}

		fmt.Println("Status:", stream.Status, " - ", stream.GetUser())
	}
}

func AddUsers(client pb.UserServiceClient) {
	reqs := []*pb.User{
		&pb.User{
			Id:    "1",
			Name:  "jho 1",
			Email: "jho1@outlook.com",
		},
		&pb.User{
			Id:    "2",
			Name:  "jho2",
			Email: "jho2@outlook.com",
		},
		&pb.User{
			Id:    "3",
			Name:  "jho3",
			Email: "jho3@outlook.com",
		},
		&pb.User{
			Id:    "4",
			Name:  "jho4",
			Email: "jho4@outlook.com",
		},
		&pb.User{
			Id:    "5",
			Name:  "jho5",
			Email: "jho5@outlook.com",
		},
	}

	stream, err := client.AddUsers(context.Background())
	if err != nil {
		log.Fatalf("Error creating request: %v", err)
	}

	for _, req := range reqs {
		stream.Send(req)
		time.Sleep(time.Second * 3)
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("Error receiving response: %v", err)
	}

	fmt.Println(res)
}
