// Code generated by mockery v2.46.3. DO NOT EDIT.

package mocks

import (
	domain "github.com/samanshahroudi/todo-service/internal/domain"
	mock "github.com/stretchr/testify/mock"
)

// SQSService is an autogenerated mock type for the SQSService type
type SQSService struct {
	mock.Mock
}

// SendMessage provides a mock function with given fields: todo
func (_m *SQSService) SendMessage(todo domain.TodoItem) error {
	ret := _m.Called(todo)

	if len(ret) == 0 {
		panic("no return value specified for SendMessage")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(domain.TodoItem) error); ok {
		r0 = rf(todo)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewSQSService creates a new instance of SQSService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewSQSService(t interface {
	mock.TestingT
	Cleanup(func())
}) *SQSService {
	mock := &SQSService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}