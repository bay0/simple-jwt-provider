// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package storage

import (
	"sync"
)

var (
	lockmigrationMockUp sync.RWMutex
)

// Ensure, that migrationMock does implement migration.
// If this is not the case, regenerate this file with moq.
var _ migration = &migrationMock{}

// migrationMock is a mock implementation of migration.
//
//     func TestSomethingThatUsesmigration(t *testing.T) {
//
//         // make and configure a mocked migration
//         mockedmigration := &migrationMock{
//             UpFunc: func() error {
// 	               panic("mock out the Up method")
//             },
//         }
//
//         // use mockedmigration in code that requires migration
//         // and then make assertions.
//
//     }
type migrationMock struct {
	// UpFunc mocks the Up method.
	UpFunc func() error

	// calls tracks calls to the methods.
	calls struct {
		// Up holds details about calls to the Up method.
		Up []struct {
		}
	}
}

// Up calls UpFunc.
func (mock *migrationMock) Up() error {
	if mock.UpFunc == nil {
		panic("migrationMock.UpFunc: method is nil but migration.Up was just called")
	}
	callInfo := struct {
	}{}
	lockmigrationMockUp.Lock()
	mock.calls.Up = append(mock.calls.Up, callInfo)
	lockmigrationMockUp.Unlock()
	return mock.UpFunc()
}

// UpCalls gets all the calls that were made to Up.
// Check the length with:
//     len(mockedmigration.UpCalls())
func (mock *migrationMock) UpCalls() []struct {
} {
	var calls []struct {
	}
	lockmigrationMockUp.RLock()
	calls = mock.calls.Up
	lockmigrationMockUp.RUnlock()
	return calls
}
