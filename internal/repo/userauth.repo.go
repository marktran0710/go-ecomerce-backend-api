package repo

type IUserAuthRepository interface {
	AddOTP(email string, otp int, expirationTime int64) error
}

type userAuthRepository struct{}

func (u *userAuthRepository) AddOTP(email string, otp int, expirationTime int64) error {
	return nil
}

func NewUserAuthRepository() IUserAuthRepository {
	return &userAuthRepository{}
}
