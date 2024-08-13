package repo

type UserRepo struct{}

type IUserRepo interface {
	GetUserByEmail(email string) bool
}

func NewUserRepo() IUserRepo {
	return &UserRepo{}
}

func (ur *UserRepo) GetUserByEmail(email string) bool {
	return true
}
