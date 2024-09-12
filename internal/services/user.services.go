package services

import (
	"github.com/marktran77/go-ecomerce-backend-api/internal/repo"
	"github.com/marktran77/go-ecomerce-backend-api/pkg/response"
)

type IUserService interface {
	Register(email string, purpose string) int
}

type UserService struct {
	userRepo repo.IUserRepo
}

func NewUserService(userRepo repo.IUserRepo) IUserService {
	return &UserService{
		userRepo: userRepo,
	}
}

// Register implements IUserService.
func (us *UserService) Register(email string, purpose string) int {
	// 0. hashEmail

	// 1. check email exists in db

	// 2. new OPT

	// 3. save OPT in Redis with expiration time

	// 4. send Email OTP

	//1. check email exists
	if us.userRepo.GetUserByEmail(email) {
		return response.ErrcodeUserHasExists
	}
	return response.ErrCodeSuccess
}
