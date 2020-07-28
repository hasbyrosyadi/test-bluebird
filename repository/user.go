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
	GetUser(id int) (*model.User, error)
}

func (u *User) GetUser(id int) (*model.User, error) {
	user := &model.User{}
	err := u.db.Where("id = ?", id).First(user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}
