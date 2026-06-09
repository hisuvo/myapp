/*
//* Database query logic।
//* Database access logic write here.
//* In future if you want to add row sql or chage to the framework and add new framework here
//* মানে database-এর সাথে সব communication এখান থেকে হবে
*/

package users

import (
	"errors"

	"gorm.io/gorm"
)

var ErrAlreadyExist = errors.New("user email already exists")

type Repository interface{
	CreateUser(u *User) error
	GetUserByEmail(email string) (*User, error)
}

// repository take db 
type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) CreateUser(user *User) error {
	result := r.db.Create(user)

	if result.Error != nil {

		if errors.Is(result.Error, gorm.ErrDuplicatedKey){
			return ErrAlreadyExist
		}

		return result.Error
	}
	
	return nil
}

func (r *repository) GetUserByEmail(email string) (*User, error) {
	var user User
	result := r.db.Where(&User{Email: email}).First(&user)

	if result.Error != nil {
		// if errors.Is(result.Error, gorm.ErrRecordNotFound){
		// 	return  nil, nil
		// }
	
		return nil, result.Error
	}
	
	return &user, nil
}
