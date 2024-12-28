package models

import (
	"fmt"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type Role string

const (
	Manager Role = "manager"
	Seller  Role = "seller"
)

type User struct {
	ID        uint      `json:"id"`
	Username  string    `json:"username"`
	Password  string    `json:"password"`
	Position  string    `json:"position"`
	Role      Role      `json:"role"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	Removed   bool      `json:"removed"`
}

func NewUser(username string, password string, position string, role Role) (*User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	u := &User{
		Username:  username,
		Password:  string(hashedPassword),
		Position:  position,
		Role:      role,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Removed:   false,
	}

	err = u.isValid()
	if err != nil {
		return nil, err
	}

	return u, nil
}

func (u *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err == nil
}

func (u *User) isValid() error {
	if len(u.Username) < 5 {
		return fmt.Errorf("username must be at least 5 characters, got %d", len(u.Username))
	}

	if len(u.Password) < 6 {
		return fmt.Errorf("password must be at least 10 characters, got %d", len(u.Password))
	}

	if u.Role == "" {
		return fmt.Errorf("role is not blank")
	}

	return nil
}
