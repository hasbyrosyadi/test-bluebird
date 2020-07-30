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

	// verifikasi data input
	if regis.Name == "" || regis.Email == "" || regis.Password == "" || regis.ConfirmPassword == "" || regis.Address == "" {
		return errors.New("Missing Parameter")
	}

	// verifikasi password
	if regis.Password != regis.ConfirmPassword {
		return errors.New("Invalid Password")
	}

	// validasi form
	newUser := model.Regis(regis)

	// insert to database user
	err := u.UserRepository.InsertUser(newUser)
	if err != nil {
		return errors.New("Internal Server Error")
	}

	return nil
}

func (u *User) Login(login *model.Login) error {

	// verifikasi data input
	if login.Email == "" || login.Password == "" {
		return errors.New("Missing Parameter")
	}

	// get user berdasarkan email
	user, err := u.UserRepository.GetUser(login.Email)
	if err != nil {
		return errors.New("Internal Server Error")
	}

	if user.Id == 0 {
		return errors.New("User Not Found")
	}

	// validasi password user dengan input
	if user.Password != login.Password {
		return errors.New("Invalid Password")
	}

	user.IsLogin = true

	// update is_login menjadi true
	err = u.UserRepository.UpdateUser(user)
	if err != nil {
		return errors.New("Internal Server Error")
	}

	return nil
}

func (u *User) Logout(logout *model.Logout) error {

	// verifikasi data input
	if logout.Email == "" {
		return errors.New("Missing Parameter")
	}

	// get user berdasarkan email
	user, err := u.UserRepository.GetUser(logout.Email)
	if err != nil {
		return errors.New("Internal Server Error")
	}

	if user.Id == 0 {
		return errors.New("User Not Found")
	}

	user.IsLogin = false

	// update is_login menjadi false
	err = u.UserRepository.UpdateUser(user)
	if err != nil {
		return errors.New("Internal Server Error")
	}

	return nil
}
