// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package fuse

import (
	"github.com/integr8ly/integration-controller/pkg/apis/integration/v1alpha1"
	"sync"
)

var (
	lockEnMasseServiceMockCreateUser sync.RWMutex
	lockEnMasseServiceMockDeleteUser sync.RWMutex
)

// EnMasseServiceMock is a mock implementation of EnMasseService.
//
//     func TestSomethingThatUsesEnMasseService(t *testing.T) {
//
//         // make and configure a mocked EnMasseService
//         mockedEnMasseService := &EnMasseServiceMock{
//             CreateUserFunc: func(userName string, realm string) (*v1alpha1.User, error) {
// 	               panic("TODO: mock out the CreateUser method")
//             },
//             DeleteUserFunc: func(userName string, realm string) error {
// 	               panic("TODO: mock out the DeleteUser method")
//             },
//         }
//
//         // TODO: use mockedEnMasseService in code that requires EnMasseService
//         //       and then make assertions.
//
//     }
type EnMasseServiceMock struct {
	// CreateUserFunc mocks the CreateUser method.
	CreateUserFunc func(userName string, realm string) (*v1alpha1.User, error)

	// DeleteUserFunc mocks the DeleteUser method.
	DeleteUserFunc func(userName string, realm string) error

	// calls tracks calls to the methods.
	calls struct {
		// CreateUser holds details about calls to the CreateUser method.
		CreateUser []struct {
			// UserName is the userName argument value.
			UserName string
			// Realm is the realm argument value.
			Realm string
		}
		// DeleteUser holds details about calls to the DeleteUser method.
		DeleteUser []struct {
			// UserName is the userName argument value.
			UserName string
			// Realm is the realm argument value.
			Realm string
		}
	}
}

// CreateUser calls CreateUserFunc.
func (mock *EnMasseServiceMock) CreateUser(userName string, realm string) (*v1alpha1.User, error) {
	if mock.CreateUserFunc == nil {
		panic("EnMasseServiceMock.CreateUserFunc: method is nil but EnMasseService.CreateUser was just called")
	}
	callInfo := struct {
		UserName string
		Realm    string
	}{
		UserName: userName,
		Realm:    realm,
	}
	lockEnMasseServiceMockCreateUser.Lock()
	mock.calls.CreateUser = append(mock.calls.CreateUser, callInfo)
	lockEnMasseServiceMockCreateUser.Unlock()
	return mock.CreateUserFunc(userName, realm)
}

// CreateUserCalls gets all the calls that were made to CreateUser.
// Check the length with:
//     len(mockedEnMasseService.CreateUserCalls())
func (mock *EnMasseServiceMock) CreateUserCalls() []struct {
	UserName string
	Realm    string
} {
	var calls []struct {
		UserName string
		Realm    string
	}
	lockEnMasseServiceMockCreateUser.RLock()
	calls = mock.calls.CreateUser
	lockEnMasseServiceMockCreateUser.RUnlock()
	return calls
}

// DeleteUser calls DeleteUserFunc.
func (mock *EnMasseServiceMock) DeleteUser(userName string, realm string) error {
	if mock.DeleteUserFunc == nil {
		panic("EnMasseServiceMock.DeleteUserFunc: method is nil but EnMasseService.DeleteUser was just called")
	}
	callInfo := struct {
		UserName string
		Realm    string
	}{
		UserName: userName,
		Realm:    realm,
	}
	lockEnMasseServiceMockDeleteUser.Lock()
	mock.calls.DeleteUser = append(mock.calls.DeleteUser, callInfo)
	lockEnMasseServiceMockDeleteUser.Unlock()
	return mock.DeleteUserFunc(userName, realm)
}

// DeleteUserCalls gets all the calls that were made to DeleteUser.
// Check the length with:
//     len(mockedEnMasseService.DeleteUserCalls())
func (mock *EnMasseServiceMock) DeleteUserCalls() []struct {
	UserName string
	Realm    string
} {
	var calls []struct {
		UserName string
		Realm    string
	}
	lockEnMasseServiceMockDeleteUser.RLock()
	calls = mock.calls.DeleteUser
	lockEnMasseServiceMockDeleteUser.RUnlock()
	return calls
}
