package services

import "github.com/marktran77/go-ecomerce-backend-api/internal/repo"

type UserService struct {
	userRepo repo.IUserRepo
}

type IUserService interface {
	GetInfoUser() string
}

func NewUserService() *UserService {
	return &UserService{
		userRepo: repo.NewUserRepo(),
	}
}

func (us *UserService) GetInfoUser() string {
	return us.userRepo.GetInfoUser()
}
