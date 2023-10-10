package postgres

import (
	"fmt"
	"project-3/stores"
)

type Conn struct {
	db string
}

func (c Conn) Create(usr stores.User) error {
	fmt.Println("user create")
	return nil
}
func (c Conn) Update(usr stores.User) error {
	fmt.Println("user update")
	return nil
}
func (c Conn) Delete(usr stores.User) error {
	fmt.Println("user delete")
	return nil
}

func New(db string) Conn {
	return Conn{db: db}
}
