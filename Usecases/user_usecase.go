package usecases

import (
	"errors"
	"blogAPI/Domain"
)

type IuserRepository interface {
	CreateUser(user *Domain.User) error
	//LoginUser(user *Domain.User) (string,string,error)
	GetUserByEamil(userID string) (Domain.User,error)
}

type JWTService interface {
	GenerateToken(User *Domain.User) (string,error)
	ValidateToken(token string) (uint, error)
}

type PasswordServices interface{
	HashPassword(password string) (string,error)
	VerifyPassword(hashedPassword string, password string) (bool,error)

}

type UserUseCase struct {
	UserRepository IuserRepository
	JWTService JWTService
	PasswordService PasswordServices
}

func NewUserUseCase(userRepository IuserRepository, jwtService JWTService, passwordService PasswordServices) Domain.IUserUseCase {
	return UserUseCase{
		UserRepository: userRepository,
		JWTService: jwtService,
		PasswordService: passwordService,
	}
}

func (usr UserUseCase) CreateUser(user *Domain.User) error {
	
	_, err := usr.UserRepository.GetUserByEamil(user.Email)
	if err == nil {
		return errors.New("email already exist")
	}

	hashedPassword, err := usr.PasswordService.HashPassword(user.Password)
	if err != nil {
		return err
	}
	user.Password = hashedPassword
	
	err = usr.UserRepository.CreateUser(user)
	if err != nil {
		return err
	}
	return nil

}

func (usr UserUseCase) LoginUser(user *Domain.User) (string,string,error) {
	existingUser, err := usr.UserRepository.GetUserByEamil(user.Email)
	if err != nil {
		return "", "", err
	}
	verified, err := usr.PasswordService.VerifyPassword(existingUser.Password, user.Password)
	if ! verified{
		return "", "", errors.New("password is incorrect")
	}
	if err != nil {
		return "", "", err
	}

	user.Id = existingUser.Id
	user.Role = existingUser.Role
	token, err := usr.JWTService.GenerateToken(user)
	return existingUser.Email, token, err
}

func (usr UserUseCase) GetUserByEamil(email string) (Domain.User,error) {
	user, err := usr.UserRepository.GetUserByEamil(email)
	if err != nil {
		return Domain.User{}, err
	}
	return user, nil
}

