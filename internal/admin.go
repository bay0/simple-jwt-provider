package internal

import (
	"errors"
	"fmt"
	"github.com/leberKleber/simple-jwt-provider/internal/storage"
	"golang.org/x/crypto/bcrypt"
)

var bcryptCost = 12
var ErrUserAlreadyExists = errors.New("user already exists")

/**
Creates new user with given email, password and claims.
return ErrUserAlreadyExists when user already exists
*/
func (p Provider) CreateUser(email, password string, claims map[string]interface{}) error {
	securedPassword, err := bcryptPassword(password)
	if err != nil {
		return fmt.Errorf("failed to bcrypt password: %w", err)
	}

	err = p.Storage.CreateUser(storage.User{
		EMail:    email,
		Password: securedPassword,
		Claims:   claims,
	})
	if err != nil {
		if errors.Is(err, storage.ErrUserAlreadyExists) {
			return ErrUserAlreadyExists
		}
		return fmt.Errorf("failed to query user with email %q: %w", email, err)
	}

	return nil
}

// DeleteUser deletes user with given email.
// return ErrUserNotFound when user does not exist
// return ErrUserStillHasTokens when user still has tokens
func (p Provider) DeleteUser(email string) error {
	err := p.Storage.DeleteUser(email)
	if err != nil {
		if errors.Is(err, storage.ErrUserNotFound) {
			return ErrUserNotFound
		} else {
			return fmt.Errorf("failed to delete user with email %q: %w", email, err)
		}
	}

	return nil
}

func bcryptPassword(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcryptCost)
}
