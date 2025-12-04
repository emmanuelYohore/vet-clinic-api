package dbmodel

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email        string `gorm:"type:varchar(255);unique;not null" json:"email"`
	Password     string `gorm:"type:varchar(255);not null" json:"password"`
	RefreshToken string `json:"refresh_token"`
}

type UserRepository interface {
	Create(user *User) (*User, error)
	FindAll() ([]*User, error)
	FindById(id uint) (*User, error)
	Update(user *User) (*User, error)
	Delete(id uint, user *User) error
	GetUserByEmail(email string) (*User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) Delete(id uint, user *User) error {
	return r.db.Delete(user, id).Error
}

func (r *userRepository) FindById(id uint) (*User, error) {
	var user User
	if err := r.db.First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil

}

func (r *userRepository) Update(user *User) (*User, error) {
	if err := r.db.Save(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (r *userRepository) Create(user *User) (*User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	user.Password = string(hashedPassword)
	if err := r.db.Create(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (r *userRepository) FindAll() ([]*User, error) {
	var users []*User
	if err := r.db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (r *userRepository) GetUserByEmail(email string) (*User, error) {
	var user User
	err := r.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, fmt.Errorf("user with email %s not found: %w", email, err)
	}
	return &user, nil
}
