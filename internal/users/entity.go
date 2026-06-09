package users

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string `json:"name" validate:"required" gorm:"type:varchar(255);not null"`
	Email    string `json:"email" validate:"required,email" gorm:"type:varchar(255);uniqueIndex;not null"`
	Password string `json:"-" validate:"required,min=6" gorm:"type:varchar(100);not null"`
}

func (u *User) hashPassword(password string) error{
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		return  err
	}

	u.Password = string(bytes)
	return nil
}

func (u *User) CheckPassword(password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password),[]byte(password))
	if err != nil {
		return err
	}
	return  nil
}