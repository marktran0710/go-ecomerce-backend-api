package services

import (
	"github.com/marktran77/go-ecomerce-backend-api/internal/repo"
	"github.com/marktran77/go-ecomerce-backend-api/pkg/response"
)

type UserService struct {
	userRepo repo.IUserRepo
}

// Register implements IUserService.
func (us *UserService) Register(email string, purpose string) int {
	//1. check email exists
	if us.userRepo.GetUserByEmail(email) {
		return response.ErrcodeUserHasExists
	}
	return response.ErrCodeSuccess
}

type IUserService interface {
	Register(email string, purpose string) int
}

func NewUserService(userRepo repo.IUserRepo) IUserService {
	return &UserService{
		userRepo: userRepo,
	}
}
