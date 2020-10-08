// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package mail

import (
	"gopkg.in/mail.v2"
	"sync"
)

var (
	lockdialerMockDial        sync.RWMutex
	lockdialerMockDialAndSend sync.RWMutex
)

// Ensure, that dialerMock does implement dialer.
// If this is not the case, regenerate this file with moq.
var _ dialer = &dialerMock{}

// dialerMock is a mock implementation of dialer.
//
//     func TestSomethingThatUsesdialer(t *testing.T) {
//
//         // make and configure a mocked dialer
//         mockeddialer := &dialerMock{
//             DialFunc: func() (mail.SendCloser, error) {
// 	               panic("mock out the Dial method")
//             },
//             DialAndSendFunc: func(msgs ...*mail.Message) error {
// 	               panic("mock out the DialAndSend method")
//             },
//         }
//
//         // use mockeddialer in code that requires dialer
//         // and then make assertions.
//
//     }
type dialerMock struct {
	// DialFunc mocks the Dial method.
	DialFunc func() (mail.SendCloser, error)

	// DialAndSendFunc mocks the DialAndSend method.
	DialAndSendFunc func(msgs ...*mail.Message) error

	// calls tracks calls to the methods.
	calls struct {
		// Dial holds details about calls to the Dial method.
		Dial []struct {
		}
		// DialAndSend holds details about calls to the DialAndSend method.
		DialAndSend []struct {
			// Msgs is the msgs argument value.
			Msgs []*mail.Message
		}
	}
}

// Dial calls DialFunc.
func (mock *dialerMock) Dial() (mail.SendCloser, error) {
	if mock.DialFunc == nil {
		panic("dialerMock.DialFunc: method is nil but dialer.Dial was just called")
	}
	callInfo := struct {
	}{}
	lockdialerMockDial.Lock()
	mock.calls.Dial = append(mock.calls.Dial, callInfo)
	lockdialerMockDial.Unlock()
	return mock.DialFunc()
}

// DialCalls gets all the calls that were made to Dial.
// Check the length with:
//     len(mockeddialer.DialCalls())
func (mock *dialerMock) DialCalls() []struct {
} {
	var calls []struct {
	}
	lockdialerMockDial.RLock()
	calls = mock.calls.Dial
	lockdialerMockDial.RUnlock()
	return calls
}

// DialAndSend calls DialAndSendFunc.
func (mock *dialerMock) DialAndSend(msgs ...*mail.Message) error {
	if mock.DialAndSendFunc == nil {
		panic("dialerMock.DialAndSendFunc: method is nil but dialer.DialAndSend was just called")
	}
	callInfo := struct {
		Msgs []*mail.Message
	}{
		Msgs: msgs,
	}
	lockdialerMockDialAndSend.Lock()
	mock.calls.DialAndSend = append(mock.calls.DialAndSend, callInfo)
	lockdialerMockDialAndSend.Unlock()
	return mock.DialAndSendFunc(msgs...)
}

// DialAndSendCalls gets all the calls that were made to DialAndSend.
// Check the length with:
//     len(mockeddialer.DialAndSendCalls())
func (mock *dialerMock) DialAndSendCalls() []struct {
	Msgs []*mail.Message
} {
	var calls []struct {
		Msgs []*mail.Message
	}
	lockdialerMockDialAndSend.RLock()
	calls = mock.calls.DialAndSend
	lockdialerMockDialAndSend.RUnlock()
	return calls
}
