package usecase

import (
	"bluebird/model"
	"bluebird/repository"
	"errors"
)

type UserUsecase interface {
	Register(regis *model.Register) error
	Login(login *model.Login) error
	Logout(logout *model.Logout) error
}

type User struct {
	UserRepository repository.UserRepository
}

func NewUser(u repository.UserRepository) UserUsecase {
	return &User{u}
}

func (u *User) Register(regis *model.Register) error {

	if regis.Name == "" || regis.Email == "" || regis.Password == "" || regis.ConfirmPassword == "" {
		return errors.New("Missing Parameter")
	}

	if regis.Password != regis.ConfirmPassword {
		return errors.New("Invalid Password")
	}

	newUser := &model.User{
		Name:     regis.Name,
		Email:    regis.Email,
		Password: regis.Password,
		Address:  regis.Address,
		Role:     "USER",
	}

	err := u.UserRepository.InsertUser(newUser)
	if err != nil {
		return errors.New("Internal Server Error")
	}

	return nil
}

func (u *User) Login(login *model.Login) error {

	if login.Email == "" || login.Password == "" {
		return errors.New("Missing Parameter")
	}

	user, err := u.UserRepository.GetUser(login.Email)
	if err != nil {
		return errors.New("Internal Server Error")
	}

	if user.Id == 0 {
		return errors.New("User Not Found")
	}

	if user.Password != login.Password {
		return errors.New("Invalid Password")
	}

	user.IsLogin = true

	err = u.UserRepository.UpdateUser(user)
	if err != nil {
		return errors.New("Internal Server Error")
	}

	return nil
}

func (u *User) Logout(logout *model.Logout) error {

	if logout.Email == "" {
		return errors.New("Missing Parameter")
	}

	user, err := u.UserRepository.GetUser(logout.Email)
	if err != nil {
		return errors.New("Internal Server Error")
	}

	if user.Id == 0 {
		return errors.New("User Not Found")
	}

	user.IsLogin = false

	err = u.UserRepository.UpdateUser(user)
	if err != nil {
		return errors.New("Internal Server Error")
	}

	return nil
}
