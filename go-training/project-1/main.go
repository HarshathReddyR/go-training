package main

import (
	"fmt"
	"go-training/project-1/models"
)

func main() {
	u1 := models.Users{
		Name:       "Abhi",
		Email:      "abhi18@gmail.com",
		Password:   "1234",
		Permission: map[string]bool{"admin": true},
	}
	fmt.Println(u1)
	for key, value := range u1.Permission {
		fmt.Println(key, value)
	}
}
