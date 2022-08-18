package model

import "time"

type User struct {
	ID             uint64    `gorm:"column:id" json:"id"`
	Email          string    `gorm:"column:email" json:"email"`
	Password       string    `gorm:"column:password" json:"password"`
	Username       string    `gorm:"column:username" json:"username"`
	NationID       string    `gorm:"nation_id" json:"nation_id"`
	GenderType     int8      `gorm:"column:gender_type" json:"gender_type"`
	BankCardNumber string    `gorm:"column:bank_card" json:"bank_card"`
	Currency       string    `gorm:"column:currency" json:"currency"`
	Balance        float64   `gorm:"column:balance" json:"balance"`
	PhoneNumber    string    `gorm:"column:phone_number" json:"phone_number"`
	City           string    `gorm:"column:city" json:"city"`
	RegisterIP     string    `gorm:"column:register_ip" json:"register_ip"`
	LastLoginIP    string    `gorm:"column:last_login_ip" json:"last_login_ip"`
	CreatedAt      time.Time `gorm:"column:created_at" json:"created_at"`
	DeletedAt      time.Time `gorm:"colum:deleted_at" json:"deleted_at"`
}

func (u *User) TableName() string {
	return "user_info"
}

type VerifyCode struct {
	ID         uint64    `gorm:"column:id" json:"id"`
	Email      string    `gorm:"column:email" json:"email"`
	Code       string    `gorm:"code" json:"code"`
	ExpireTime time.Time `gorm:"expire_time" json:"expire_time"`
	CreatedAt  time.Time `gorm:"column:created_at" json:"created_at"`
	DeletedAt  time.Time `gorm:"column:deleted_at" json:"deleted_at"`
}

func (vc *VerifyCode) TableName() string {
	return "verify_code"
}
