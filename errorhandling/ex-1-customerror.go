package main

import (
	"errors"
	"fmt"
)

var ErrStringIsEmpty = errors.New("String is Empty")

func FetchData(s string) (string, error) {
	if s == "" {
		return "", ErrStringIsEmpty
	}
	return s, nil
}
func main() {
	s, err := FetchData("hi")
	fmt.Println(s, err)
}
