package repo

type UserRepo struct{}

type IUserRepo interface {
	GetInfoUser() string
}

func NewUserRepo() *UserRepo {
	return &UserRepo{}
}

func (ur *UserRepo) GetInfoUser() string {
	return "HauTran"
}
