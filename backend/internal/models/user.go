package models

import (
	"fmt"
	"regexp"
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
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	Position  string    `json:"position"`
	Role      Role      `json:"role"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	Removed   bool      `json:"removed"`
}

func NewUser(email string, password string, position string, role Role) (*User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	u := &User{
		Email:     email,
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
	regexEmail := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)

	if !regexEmail.MatchString(u.Email) {
		return fmt.Errorf("formato de e-mail inválido: %s", u.Email)
	}

	if len(u.Password) < 6 {
		return fmt.Errorf("a senha não pode ter um tamanho menor que 6 %d", len(u.Password))
	}

	if u.Role == "" {
		return fmt.Errorf("a permissão não pode ser vazia")
	}

	return nil
}
