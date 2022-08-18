package main

import (
	"context"
	"io"
	"log"

	"github.com/shaineminkyaw/grpc_lab/login/config"
	pb "github.com/shaineminkyaw/grpc_lab/login/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	//
	conn, err := grpc.Dial(config.GRPC.AuthGrpc, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("error on connected to login server %v", err)
	}

	loginClient := pb.NewLoginServiceClient(conn)
	SendIDtoLoginServer(loginClient)
}

func SendIDtoLoginServer(client pb.LoginServiceClient) {
	//
	ids := make([]*pb.RequestID, 0)
	for i := 1; i < 6; i++ {
		ids = append(ids, &pb.RequestID{
			Id: uint64(i),
		})
	}
	allID := &pb.LoginRequest{
		RequestId: ids,
	}

	stream, err := client.Login(context.Background(), allID)
	if err != nil {
		log.Fatalf("error on reading data %v", err)
	}
	var i int64 = 1
	for {
		token, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("error on receiving data %v", err)
		}
		log.Printf("%v : %v", i, token.AllToken)
		i++
	}

}
