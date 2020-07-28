package repository

import (
	"bluebird/model"

	"github.com/jinzhu/gorm"
)

type User struct {
	db *gorm.DB
}

func NewUser(d *gorm.DB) UserRepository {
	return &User{db: d}
}

type UserRepository interface {
	GetUser(email string) (*model.User, error)
	InsertUser(user *model.User) error
	UpdateUser(user *model.User) error
}

func (u *User) GetUser(email string) (*model.User, error) {
	user := &model.User{}
	err := u.db.Where("email = ?", email).First(user).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return user, nil
}

func (u *User) InsertUser(user *model.User) error {
	if err := u.db.Create(user).Error; err != nil {
		return err
	}
	return nil
}

func (u *User) UpdateUser(user *model.User) error {
	if err := u.db.Save(&user).Error; err != nil {
		return err
	}
	return nil
}
