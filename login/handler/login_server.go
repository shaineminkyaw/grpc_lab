package handler

import (
	"log"
	"net"

	"github.com/shaineminkyaw/grpc_lab/login/config"
	"github.com/shaineminkyaw/grpc_lab/login/ds"
	"github.com/shaineminkyaw/grpc_lab/login/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type LoginServer struct {
	//
	pb.UnimplementedLoginServiceServer
	Database *ds.DataSource
}

func NewLoginServer() (*LoginServer, error) {
	db := ds.LoginConnectDB()
	return &LoginServer{
		Database: db,
	}, nil
}

func RunLoginGRPCServer() {
	source_server, err := NewLoginServer()
	if err != nil {
		log.Fatalf("error on source server %v", err)
	}

	grpcServer := grpc.NewServer()

	pb.RegisterLoginServiceServer(grpcServer, source_server)
	reflection.Register(grpcServer)

	lis, err := net.Listen("tcp", "localhost:50052")
	if err != nil {
		log.Fatalf("error on listening %v", err)
	}

	log.Printf("starting GRPC server %v", config.Init().GRPCAddress)
	err = grpcServer.Serve(lis)
	if err != nil {
		log.Fatalf("error on GRPC server %v", err)
	}
}
