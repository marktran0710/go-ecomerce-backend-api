package repo

import (
	"fmt"
	"time"

	"github.com/marktran77/go-ecomerce-backend-api/global"
)

type IUserAuthRepo interface {
	AddOTP(email string, otp int, expirationTime int64) error
}

type userAuthRepository struct{}

func (u *userAuthRepository) AddOTP(email string, otp int, expirationTime int64) error {
	key := fmt.Sprintf("user:%s:otp", email) // user:email:otp
	return global.Rdb.SetEx(ctx, key, otp, time.Duration(expirationTime)).Err()
}

func NewUserAuthRepository() IUserAuthRepo {
	return &userAuthRepository{}
}
