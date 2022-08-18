package handler

import (
	"log"
	"net"

	"github.com/shaineminkyaw/grpc_lab/user_info/config"
	"github.com/shaineminkyaw/grpc_lab/user_info/ds"
	"github.com/shaineminkyaw/grpc_lab/user_info/pb"
	"github.com/shaineminkyaw/grpc_lab/user_info/repository"
	"github.com/shaineminkyaw/grpc_lab/user_info/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type UserServer struct {
	pb.UnimplementedUserServiceServer
	UserService service.UserService
	Database    *ds.DataSource
}

func NewUserGrpcServer() (*UserServer, error) {
	db := ds.UserConnectDB()
	user_repo := repository.NewUserRepo(db)
	user_service := service.NewUserService(user_repo)
	return &UserServer{
		UserService: user_service,
		Database:    db,
	}, nil
}

func RunUserGrpcServer() {
	//

	souce_server, err := NewUserGrpcServer()
	if err != nil {
		log.Fatalf("error on source server %v", err)
	}

	grpcServer := grpc.NewServer()

	pb.RegisterUserServiceServer(grpcServer, souce_server.UnimplementedUserServiceServer)
	reflection.Register(grpcServer)

	lis, err := net.Listen("tcp", config.GRPC.AuthGrpc)
	if err != nil {
		log.Fatalf("error on not working listen %v", err)
	}

	log.Printf("starting GRPC server %v", config.GRPC.AuthGrpc)

	err = grpcServer.Serve(lis)
	if err != nil {
		log.Fatalf("error on grpc server not working %v", err)
	}
}
