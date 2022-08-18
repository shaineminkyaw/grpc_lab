package main

import (
	"context"
	"log"

	pb "github.com/shaineminkyaw/grpc_lab/login/pb"
	"google.golang.org/grpc"
)

func main() {
	//
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())
	conn, err := grpc.Dial("localhost:50052", opts...)
	defer conn.Close()
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
	var i int64 = 0
	resp, err := stream.Recv()
	if err != nil {
		log.Fatalf("error on receiving data %v", err)
	}

	for _, item := range resp.AllToken {
		i++
		log.Printf("%v......%v\n", i, item.Token)

	}

}
