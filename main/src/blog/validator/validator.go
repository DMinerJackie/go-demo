package main

import (
	"fmt"
	"gopkg.in/go-playground/validator.v9"
)

// User contains user information
type UserInfo struct {
	FirstName string `validate:"required"`
	LastName  string `validate:"required"`
	Age       uint8  `validate:"gte=0,lte=100"`
	Email     string `validate:"required,email"`
}

func main() {
	validate := validator.New()
	user := &UserInfo{
		FirstName: "Badger",
		LastName:  "Smith",
		Age:       105,
		Email:     "",
	}
	err := validate.Struct(user)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			fmt.Println(err)
		}
		return
	}
	fmt.Println("success")
}
