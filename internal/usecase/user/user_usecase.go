package usecase

import (
	model "king/internal/domain/model/user"
	repository "king/internal/domain/repository/user"

	"golang.org/x/crypto/bcrypt"
)


type UserUsecase struct {
	userRepo repository.UserRepository
}

func NewUserRepository(userRepo repository.UserRepository) *UserUsecase {
	return &UserUsecase{userRepo: userRepo}
}

func (u *UserUsecase) RegisterUser(username, password string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user := &model.User{
		Username: username,
		Password: string(hashedPassword),
	}

	return u.userRepo.CreateUser(user)

}

func(u *UserUsecase) GetUserByUsername(username string) (*model.User, error) {

	user, err := u.userRepo.GetUserByUsername(username)
	if err != nil {
		return nil, err
	}

	return user, nil

}