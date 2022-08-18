package handler

// import (
// 	"log"

// 	"github.com/shaineminkyaw/grpc_lab/user_info/model"
// 	"github.com/shaineminkyaw/grpc_lab/user_info/pb"
// )

// func (server *UserServer) FilterUserFromLoginserver(req *pb.AllIDRequest, stream pb.UserService_FilterUserFromLoginserverServer) error {
// 	//

// 	for _, id := range req.AllId {
// 		user := &model.User{}
// 		var cond bool = true
// 		err := server.Database.Data.Model(&model.User{}).Where("id = ?", id.Id).First(&user).Error
// 		if err != nil {
// 			log.Fatalf("error on find user %v", err)
// 		}
// 		if user != nil {
// 			cond = true
// 		}
// 	}

// }
