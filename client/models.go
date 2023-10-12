package main

import (
	"time"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID        uint   `json:"id" gorm:"primaryKey;unique"  `
	FirstName string `json:"first_name" validate:"required,min=2,max=100"`
	LastName  string `json:"last_name"  validate:"required,min=2,max=100"`
	Email     string `json:"email" validate:"email,required" `
	Password  string `json:"password" validate:"required,min=6"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type SignUpRequest struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

type ErrResponse struct{
	Error string `json:"error"`
	error string
}

type SuccResponse struct{
	Sucess string `json:"success"`
}

func (u *User) HashPassword(pass string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(pass), 10)
	if err != nil {
		return err
	}

	u.Password = string(bytes)
	return nil
}

func (u *User) ComparePassword(incoming string, pass User) error {

	if err := bcrypt.CompareHashAndPassword([]byte(incoming), []byte(u.Password)); err != nil {
		return err
	}
	return nil
}
