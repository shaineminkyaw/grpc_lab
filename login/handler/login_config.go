package handler

import (
	"context"
	"io"
	"log"

	pb "github.com/shaineminkyaw/grpc_lab/login/pb"
	pb1 "github.com/shaineminkyaw/grpc_lab/user_info/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func (server *LoginServer) Login(req *pb.LoginRequest, stream pb.LoginService_LoginServer) error {
	//
	ids := make([]uint64, 0)
	tokens := make([]*pb.ResponseToken, 0)
	for _, id := range req.RequestId {
		ids = append(ids, id.Id)
	}
	allBool, err := CheckUserToUserServer(ids)
	if err != nil {
		log.Fatalf("error on return data %v", err)
	}
	for _, universe := range allBool {
		if universe {
			//
			abc := "asfdashfgdhasfdj"
			tokens = append(tokens, &pb.ResponseToken{
				Token: abc,
			})
		} else {
			//
			abc := ""
			tokens = append(tokens, &pb.ResponseToken{
				Token: abc,
			})
		}
	}
	stream.Send(&pb.LoginResponse{
		AllToken: tokens,
	})
	return nil

}

func CheckUserToUserServer(req []uint64) ([]bool, error) {
	//

	conn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("error on connection to User Server %v", err)
	}

	request := make([]*pb1.LoginServerRequest, 0)

	for _, id := range req {
		request = append(request, &pb1.LoginServerRequest{
			Id: id,
		})
	}

	allID := &pb1.AllIDRequest{
		AllId: request,
	}
	client := pb1.NewUserServiceClient(conn)
	stream, err := client.FilterUserFromLoginserver(context.Background(), allID)
	if err != nil {
		log.Fatalf("error on reading data %v", err)
	}

	allBool := make([]bool, 0)

	for {
		bools, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("error on receiving data %v", err)
		}
		allBool = append(allBool, bools.Result)

	}
	return allBool, nil

}
