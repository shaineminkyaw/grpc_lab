package handler

import (
	"github.com/shaineminkyaw/grpc_lab/login/ds"
	"github.com/shaineminkyaw/grpc_lab/login/pb"
)

type LoginServer struct {
	//
	pb.UnimplementedLoginServiceServer
	Database ds.DataSource
}
