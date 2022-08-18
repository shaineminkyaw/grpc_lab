package model

import "time"

type UserToken struct {
	Id                     uint64    `gorm:"column:id" json:"id"`
	Uid                    uint64    `gorm:"column:uid" json:"uid"`
	AccessToken            string    `gorm:"access_token" json:"access_token"`
	AccessTokenExpireTime  time.Time `gorm:"accessToken_expire_time" json:"accessToken_expire_time"`
	RefreshTokenID         string    `gorm:"column:refresh_token_id" json:"refresh_token_id"`
	RefreshToken           string    `gorm:"column:refresh_token" json:"refresh_token"`
	RefreshTokenExpireTime time.Time `gorm:"column:refreshToken_expire_time" json:"refreshToken_expire_time"`
	CreatedAt              time.Time `gorm:"column:creaed_at" json:"created_at"`
	DeletedAt              time.Time `gorm:"column:deleted_at" json:"deleted_at"`
}

func (ut *UserToken) TableName() string {
	return "user_token"
}
