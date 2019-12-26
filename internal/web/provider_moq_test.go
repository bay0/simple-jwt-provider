// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package web

import (
	"sync"
)

var (
	lockProviderMockCreateUser sync.RWMutex
	lockProviderMockLogin      sync.RWMutex
)

// Ensure, that ProviderMock does implement Provider.
// If this is not the case, regenerate this file with moq.
var _ Provider = &ProviderMock{}

// ProviderMock is a mock implementation of Provider.
//
//     func TestSomethingThatUsesProvider(t *testing.T) {
//
//         // make and configure a mocked Provider
//         mockedProvider := &ProviderMock{
//             CreateUserFunc: func(email string, password string) error {
// 	               panic("mock out the CreateUser method")
//             },
//             LoginFunc: func(email string, password string) (string, error) {
// 	               panic("mock out the Login method")
//             },
//         }
//
//         // use mockedProvider in code that requires Provider
//         // and then make assertions.
//
//     }
type ProviderMock struct {
	// CreateUserFunc mocks the CreateUser method.
	CreateUserFunc func(email string, password string) error

	// LoginFunc mocks the Login method.
	LoginFunc func(email string, password string) (string, error)

	// calls tracks calls to the methods.
	calls struct {
		// CreateUser holds details about calls to the CreateUser method.
		CreateUser []struct {
			// Email is the email argument value.
			Email string
			// Password is the password argument value.
			Password string
		}
		// Login holds details about calls to the Login method.
		Login []struct {
			// Email is the email argument value.
			Email string
			// Password is the password argument value.
			Password string
		}
	}
}

// CreateUser calls CreateUserFunc.
func (mock *ProviderMock) CreateUser(email string, password string) error {
	if mock.CreateUserFunc == nil {
		panic("ProviderMock.CreateUserFunc: method is nil but Provider.CreateUser was just called")
	}
	callInfo := struct {
		Email    string
		Password string
	}{
		Email:    email,
		Password: password,
	}
	lockProviderMockCreateUser.Lock()
	mock.calls.CreateUser = append(mock.calls.CreateUser, callInfo)
	lockProviderMockCreateUser.Unlock()
	return mock.CreateUserFunc(email, password)
}

// CreateUserCalls gets all the calls that were made to CreateUser.
// Check the length with:
//     len(mockedProvider.CreateUserCalls())
func (mock *ProviderMock) CreateUserCalls() []struct {
	Email    string
	Password string
} {
	var calls []struct {
		Email    string
		Password string
	}
	lockProviderMockCreateUser.RLock()
	calls = mock.calls.CreateUser
	lockProviderMockCreateUser.RUnlock()
	return calls
}

// Login calls LoginFunc.
func (mock *ProviderMock) Login(email string, password string) (string, error) {
	if mock.LoginFunc == nil {
		panic("ProviderMock.LoginFunc: method is nil but Provider.Login was just called")
	}
	callInfo := struct {
		Email    string
		Password string
	}{
		Email:    email,
		Password: password,
	}
	lockProviderMockLogin.Lock()
	mock.calls.Login = append(mock.calls.Login, callInfo)
	lockProviderMockLogin.Unlock()
	return mock.LoginFunc(email, password)
}

// LoginCalls gets all the calls that were made to Login.
// Check the length with:
//     len(mockedProvider.LoginCalls())
func (mock *ProviderMock) LoginCalls() []struct {
	Email    string
	Password string
} {
	var calls []struct {
		Email    string
		Password string
	}
	lockProviderMockLogin.RLock()
	calls = mock.calls.Login
	lockProviderMockLogin.RUnlock()
	return calls
}
