package ds

import (
	"fmt"
	"log"

	"github.com/shaineminkyaw/grpc_lab/user_info/config"
	"github.com/shaineminkyaw/grpc_lab/user_info/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DataSource struct {
	Data *gorm.DB
}

var User_DB *gorm.DB

func UserConnectDB() *DataSource {
	//
	conf := config.Init()
	host := conf.Host
	port := conf.Port
	dbName := conf.DB
	dbUser := conf.DBUser
	dbPassword := conf.DBPassword

	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local", dbUser, dbPassword, host, port, dbName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("error on connecting to database!")
	} else {
		log.Printf("Connected Database :::")
	}

	User_DB = db
	db.AutoMigrate(
		&model.User{},
		&model.VerifyCode{},
		&model.UserBankCard{},
		&model.UserCityTotalBankCard{},
	)

	return &DataSource{
		Data: db,
	}
}
