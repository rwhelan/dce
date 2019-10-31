// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import (
	api "github.com/Optum/Redbox/pkg/api"
	events "github.com/aws/aws-lambda-go/events"

	mock "github.com/stretchr/testify/mock"
)

// UserDetailer is an autogenerated mock type for the UserDetailer type
type UserDetailer struct {
	mock.Mock
}

// GetUser provides a mock function with given fields: event
func (_m *UserDetailer) GetUser(event *events.APIGatewayProxyRequest) *api.User {
	ret := _m.Called(event)

	var r0 *api.User
	if rf, ok := ret.Get(0).(func(*events.APIGatewayProxyRequest) *api.User); ok {
		r0 = rf(event)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*api.User)
		}
	}

	return r0
}
