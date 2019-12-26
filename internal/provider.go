package internal

import (
	"github.com/leberKleber/simple-jwt-provider/internal/storage"
)

//go:generate moq -out storage_moq_test.go . Storage
type Storage interface {
	User(email string) (storage.User, error)
	CreateUser(user storage.User) error
}

//go:generate moq -out jwt_generator_moq_test.go . JWTGenerator
type JWTGenerator interface {
	Generate(email string) (string, error)
}

type Provider struct {
	Storage      Storage
	JWTGenerator JWTGenerator
}
