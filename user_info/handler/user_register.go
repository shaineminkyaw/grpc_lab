package handler

import (
	"context"
	"log"
	"time"

	"github.com/mazen160/go-random"
	"github.com/shaineminkyaw/grpc_lab/user_info/model"
	"github.com/shaineminkyaw/grpc_lab/user_info/pb"
	util "github.com/shaineminkyaw/grpc_lab/user_info/utils"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
	"gorm.io/gorm"
)

func (server *UserServer) CreateUser(ctx context.Context, req *pb.UserRequest) (*pb.UserResponse, error) {

	vUser := &model.VerifyCode{}
	//@@@validate verifycode
	err := server.Database.Data.Model(&model.VerifyCode{}).Where("email = ?", req.GetEmail()).First(&vUser).Error
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error on validate verify code ")
	}
	if vUser == nil {
		return nil, status.Errorf(codes.NotFound, "email not registered for verify code !")
	}
	if req.GetCode() != vUser.Code || time.Now().Unix() > vUser.ExpireTime.UnixNano() {
		return nil, status.Errorf(codes.Internal, "verifycode invalid")
	}

	//hash password
	userPassword, err := util.HashPassword(req.GetPassword())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error on hash password")
	}
	charset := "12345678"
	length := 8
	usr, _ := random.Random(length, charset, true)
	userName := "U" + usr
	currency := "USD"
	bankCard, err := util.GetBankCardNumber(req.City)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "erorr on generate bank card")
	}
	user := &model.User{
		Username:       userName,
		Password:       userPassword,
		Email:          req.GetEmail(),
		RegisterIP:     "",
		LastLoginIP:    "",
		NationID:       req.GetNationId(),
		BankCardNumber: bankCard,
		City:           req.GetCity(),
		Balance:        100,
		Currency:       currency,
		GenderType:     int8(req.GetGenderType()),
	}

	fUser := &model.User{}
	data := server.Database.Data.Model(&model.User{}).Where("email =?", req.GetEmail())
	err = data.First(&fUser).Error
	if ctx.Err() == context.Canceled {
		log.Printf("request is canceled")
		return nil, status.Errorf(codes.Aborted, "request is Cancel")
	}
	if err == gorm.ErrRecordNotFound {
		err := server.Database.Data.Model(&model.User{}).Create(&user).Error
		if err != nil {
			return nil, status.Errorf(codes.Internal, "erorr on user create !")
		}
		err = util.SaveUserBankCard(req.GetCity(), user.ID, bankCard)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "erorr on store bank card ")
		}
	} else {
		return nil, status.Errorf(codes.Internal, "user already exists")
	}

	resp := &pb.UserResponse{
		User: convert(user),
	}

	return resp, nil
}

func convert(user *model.User) *pb.User {
	//
	return &pb.User{
		Id:          user.ID,
		Email:       user.Email,
		Username:    user.Username,
		Password:    user.Password,
		PhoneNumber: user.PhoneNumber,
		NationId:    user.NationID,
		City:        user.City,
		Bankcard:    user.BankCardNumber,
		Currency:    user.Currency,
		Balance:     user.Balance,
		GenderType:  int32(user.GenderType),
		RegisterIp:  user.RegisterIP,
		LastLoginIp: user.LastLoginIP,
		CreatedAt:   timestamppb.New(user.CreatedAt),
		DeletedAt:   timestamppb.New(user.DeletedAt),
	}
}
