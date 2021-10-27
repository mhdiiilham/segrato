package password

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

type service struct{}

func NewService() *service {
	return &service{}
}

func (s *service) HashPassword(password string) (string, error) {
	p, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil {
		fmt.Println("err", err)
		return "", err
	}

	return string(p), nil
}
