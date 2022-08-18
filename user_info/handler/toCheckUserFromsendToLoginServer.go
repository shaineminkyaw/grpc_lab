package handler

import (
	"github.com/shaineminkyaw/grpc_lab/user_info/model"
	"github.com/shaineminkyaw/grpc_lab/user_info/pb"
	"gorm.io/gorm"
)

func (server *UserServer) FilterUserFromLoginserver(req *pb.AllIDRequest, stream pb.UserService_FilterUserFromLoginserverServer) error {
	//
	user := &model.User{}
	for _, allID := range req.GetAllId() {
		err := server.Database.Data.Model(&model.User{}).Where("id = ?", allID.Id).First(&user).Error
		if err == gorm.ErrRecordNotFound {
			stream.Send(&pb.ResponseToLoginServer{
				Result: false,
			})
		} else {
			stream.Send(&pb.ResponseToLoginServer{
				Result: true,
			})
		}
	}
	return nil

}
