package models

type Users struct {
	Name       string
	Email      string
	Password   string
	Permission map[string]bool
}
